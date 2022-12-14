import os
import json
from dotenv import load_dotenv
from repository import Repository
from logic import Logic

load_dotenv("env/.env")

def lambda_handler(event, context):
    repo = Repository(os.environ.get("BUCKET_NAME"))
    df = repo.get_dataframe(event["Key"])
    idx = Logic.extract_target_index(df, event["Method"])
    if idx == None:
        return {"statusCode": 400, "body": "invalid method"}    
    body = Logic.get_record_body(df, idx)

    return {
        "statusCode": 200,
        "key": event["Key"],
        "method": event["Method"],
        "body": json.dumps(body),
    }
