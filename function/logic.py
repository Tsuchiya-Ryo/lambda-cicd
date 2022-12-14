import typing as t
from pandas import DataFrame

class Logic(object):

    @staticmethod
    def extract_target_index(df: DataFrame, method: str) -> t.Optional[int]:
        match method:
            case "max":
                return df["number"].idxmax()
            case "min":
                return df["number"].idxmin()
            case _:
                return None

    @staticmethod
    def get_record_body(df: DataFrame, idx: int) -> t.Dict[str, str]:
        body = {}
        for col in df.columns:
            body[col] = str(df.loc[idx, col])
        return body
