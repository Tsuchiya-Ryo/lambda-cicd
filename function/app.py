import json
from dotenv import load_dotenv
from repository import Repository

load_dotenv("env/.env")

def lambda_handler(event, context):
    repo = Repository(BUCKET_NAME)
    df = repo.get_dataframe(event["Key"])
    if event["Method"] == "max":
        idx = df["number"].idxmax()
    elif event["Method"] == "min":
        idx = df["number"].idxmin()
    else:
        return {"statusCode": 400, "body": "invalid method"}

    body = {}
    for col in df.columns:
        body[col] = str(df.loc[idx, col])

    return {
        "statusCode": 200,
        "key": event["Key"],
        "method": event["Method"],
        "body": json.dumps(body),
    }
