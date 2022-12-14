FROM python:3.10-slim

COPY . /function
WORKDIR /function

RUN pip3 install --upgrade pip && \
    pip3 install --no-cache-dir --target ./ awslambdaric && \
    pip3 install --no-cache-dir -r requirements.txt

ENTRYPOINT [ "/usr/local/bin/python", "-m", "awslambdaric" ]
CMD [ "app.lambda_handler" ]
