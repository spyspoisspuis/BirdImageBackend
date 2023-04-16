
## Installation (through docker-compose)

1. Copy `.env.template` and rename it to `.env` and change the configurations as wish.
2. Choose whether to initialize and dump mariadb with full test data or only DDL & User
    * If test results data is needed (all table inserts), append `.temp` to the `000_init-wo-test.sql` and delete the `.temp` from the `000_init-w-test.sql.temp`
    * If only required insertion (excluding test result inserts) is needed, leave the folder `maria-db` as is.
3. In case of first-time deployment (in which SQL script will be initialized)
```bash
docker-compose up
```
4. If the source code is edited, only restarting the edited container is required. (If the containers are still running)
```bash
# Restarting Web Backend
docker-compose restart web
```
5. If the source code is edited while the containers are down, flag `--build` is needed to bring up the services with new code
```bash
docker-compose up --build
```
6. If you need a fresh reset (like deleting all of the previous kept data), remove the folder `/volumes` and bring up the services again.
```bash
# Delete the /volumes (/volumes/database if you only need to reset the db)
rm -rf ./volumes
docker-compose up --build
```
