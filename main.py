import os
from dotenv import load_dotenv

import json
import requests


class GetDailyIntegrals:
    def __init__(self):
        self.headers = {
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
        }
        self.diffs = ["BEGINNER", "EASY", "MEDIUM", "HARD"]

        self.api_endpoint_a = str(os.getenv("API_ENDPOINT_A"))
        self.api_endpoint_b = str(os.getenv("API_ENDPOINT_B"))

        self.path = "./data/"

    def get_day(self):
        r = requests.post(
            url=self.api_endpoint_a,
            json={},
            headers=self.headers,
        )
        if r.status_code != 200:
            return 0
        return int(r.json()["day"])

    def create_payload(self, diff):
        day = self.get_day()
        if day == 0:
            return {}
        payload = {"type": "integral", "day": day, "difficulty": diff}
        return payload

    def get_integral(self, payload):
        r = requests.post(
            url=self.api_endpoint_b,
            json=payload,
            headers=self.headers,
        )

        if r.status_code != 200:
            return {}

        json = r.json()["puzzle"]
        return {
            "day": json["day"],
            "difficulty": json["difficulty"],
            "title": json["title"],
            "latex": json["latex"],
        }

    def get_and_save(self):
        intgs_data = []

        for diff in self.diffs:
            print(f"getting {diff} integral")
            p = self.create_payload(diff)

            if p == {}:
                print("failed to get integral")
                continue

            intg_json = self.get_integral(p)
            intgs_data.append(intg_json)

        # file_name = self.path + "di" + str(intgs_data[0]["day"]) + ".json"
        file_name = self.path + "dailyintegral.json"
        with open(file_name, "w") as f:
            json.dump({"puzzles": intgs_data}, f)
        print(f"data saved to {file_name}")


load_dotenv()

gti = GetDailyIntegrals()
gti.get_and_save()
