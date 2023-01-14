#! /bin/bash

docker login --username AWS --password $(aws ecr get-login-password --region ${AWS_REGION}) ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com
cd ${FUNCTION_DIR}
docker build -t ${REPOSITORY_NAME} .
docker tag ${REPOSITORY_NAME}:latest ${REPOSITORY_URL}:latest
docker push ${REPOSITORY_URL}:latest
