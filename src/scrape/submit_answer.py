from datetime import datetime, date, timezone
import os
from dotenv import load_dotenv

import json
import requests


def submit_answer(diff, ans):
    api_endpoint_c = str(os.getenv("API_ENDPOINT_C"))

    headers = {
        "Apikey": str(os.getenv("API_KEY")),
        "Authorization": "Bearer " + str(os.getenv("AUTHORIZATION")),
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
    }

    user_id = str(os.getenv("USER_ID"))

    date_ = date.today().isoformat()
    now = datetime.now(timezone.utc)
    timestamp = now.strftime("%Y-%m-%dT%H:%M:%S.") + f"{now.microsecond // 1000:03d}Z"

    submit_payload = json.dumps(
        {
            "user_id": user_id,
            "date": date_,
            "calculus_type": "integral",
            "difficulty": diff,
            "attempts_used": 1,
            "hints_used": 0,
            "elapsed_ms": 13293,
            "last_answer": [ans],
            "status": "in_progress",
            "working_data": None,
            "last_fluid_latex": None,
            "updated_at": timestamp,
        }
    )

    print(submit_payload)

    r = requests.post(
        url=api_endpoint_c,
        json=submit_payload,
        headers=headers,
    )

    print(r.status_code)
    print(r.text)

    if r.status_code != 200 and r.status_code != 201:
        print("Could not submit answer")
        return


load_dotenv()
submit_answer("easy", "22")
