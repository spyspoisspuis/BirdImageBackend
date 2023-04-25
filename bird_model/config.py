import torch

# Path to the PyTorch model file
MODEL_PATH = 'model/birdClassifier.pt'

# Load the PyTorch model
model = torch.load(MODEL_PATH, map_location=torch.device('cpu'))
model.eval()

# Path to the folder where uploaded images are stored
UPLOAD_FOLDER = 'uploads'

# Allowed image extensions
ALLOWED_EXTENSIONS = {'jpg', 'jpeg', 'png'}

# Maximum image size (in bytes)
MAX_CONTENT_LENGTH = 16 * 1024 * 1024  # 16MB
