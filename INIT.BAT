@ECHO OFF
ECHO CREANDO DIRECCTORIO DE DATOS
IF NOT EXIST db_data\  MKDIR db_data
docker compose up -d --force-recreate