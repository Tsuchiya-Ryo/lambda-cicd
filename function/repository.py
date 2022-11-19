import boto3
import pandas as pd
from io import StringIO
from constants import ENCODING, DELIMITER

class Repository(object):
    def __init__(self, bucket_name: str):
        self.client = boto3.client('s3')
        self.bucket_name = bucket_name

    def get_dataframe(self, key: str) -> pd.DataFrame:
        body = self.client.get_object(Bucket=self.bucket_name, Key=key)["Body"].read()
        return pd.read_csv(StringIO(body.decode(ENCODING)), delimiter=DELIMITER)
