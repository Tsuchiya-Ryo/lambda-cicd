import pandas as pd
from unittest import TestCase
from function.logic import Logic

class TestLogic(TestCase):
    def setUp(self) -> None:
        self.df = pd.DataFrame({
            "date": ["2022-01-01", "2022-01-02", "2022-01-03", "2022-01-04", "2022-01-05"],
            "name": ["Prof. Emmanuelle", "Queen Ashlee", "Dr. Angeline ", "Dr. Angelica", "Prof. Etha"],
            "id": ["c04d7a84687e4", "dfa90fbbee9c", "915878230248", "9fcfc821c7fe", "71fe97aa97d8"],
            "age": [45, 59, 81, 47, 78],
            "number": [60.0, 70.0, 80.0, 90.0, 100.0]
        })

    def test_extract_target_index(self):
        res = Logic.extract_target_index(self.df, "max")
        assert res == 4

    def test_extract_target_invalid_method(self):
        res = Logic.extract_target_index(self.df, "invalid_method")
        assert res == None
