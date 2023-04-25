from flask import Flask, request
from flask_cors import CORS
from efficientnet_pytorch import EfficientNet
from torchvision import transforms,utils
from datetime import datetime
import os
import numpy as np
from PIL import Image
import torch
import logging
import csv

app = Flask(__name__)
CORS(app)
app.config['UPLOAD_FOLDER'] = '/uploads'
#logging initialize
# create a logger
logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)

# create a file handler
handler = logging.FileHandler('app.log')
handler.setLevel(logging.DEBUG)

# create a formatter and add it to the handler
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)

# add the handler to the logger
logger.addHandler(handler)
#end logging

device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')

# Define the model
model = EfficientNet.from_name('efficientnet-b7')
num_classes = 515
in_features = model._fc.in_features
model.classifiers = torch.nn.Linear(in_features, num_classes)
model = model.double() # Convert model to double precision
model.to(device)

# Load the pre-trained state dictionary and create a new state dictionary with renamed keys
state_dict = torch.load('birdClassifier.pt', map_location=torch.device('cpu'))
new_state_dict = {}
for k, v in state_dict.items():
    if 'classifier' in k:
        new_k = k.replace('classifier', 'classifiers')
        new_state_dict[new_k] = v
    else:
        new_state_dict[k] = v

# Load the new state dictionary into the model
model.load_state_dict(new_state_dict)
model.eval()


if not os.path.exists(app.config['UPLOAD_FOLDER']):
    os.makedirs(app.config['UPLOAD_FOLDER'])

rows = []
with open('classname.csv','r') as file :
    csv_reader = csv.reader(file)
    rows = list(csv_reader)

classname = {row[0]:row[1] for row in rows}
print(classname)
#pre-process image
test_transforms2 = transforms.Compose([
    transforms.Resize(256),
    transforms.CenterCrop(224),
    transforms.ToTensor(),
])



@app.route("/")
def homepage():
    return "Hello world"


@app.route('/predict', methods=['POST'])
def classifier():
    logger.debug(request.__dict__)
    try : 
        # Load image
        image_file = request.files['image']
        now = datetime.now()
        current_time = now.strftime("%Y-%m-%d %H:%M:%S")   
        filename =  current_time+"_"+image_file.filename
        image_path = os.path.join(app.config['UPLOAD_FOLDER'], filename)
        image_file.save(image_path)

        with torch.no_grad():
            # Preprocess image
            image = Image.open(image_path).convert('RGB')
            image_tensor = test_transforms2(image).unsqueeze(0).to(device)
            image_tensor = image_tensor.double()

            # Run inference
            output = model(image_tensor)
            # Get predicted class index
            _, predicted = torch.max(output.data, 1)
            loggingMessage = filename + " result: " + str(predicted[0].item())
            logger.debug(loggingMessage)
            # Return prediction as JSON
            i = str(predicted[0].item()+1)
            bird_id = "B" + ("0" * (3-len(i))) + i
            return classname[bird_id]
        
    except Exception as e :
        logger.error(str(e))
        return {"error": "An error occurred during prediction."}, 500
    
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5100)

