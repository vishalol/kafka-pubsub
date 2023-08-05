import requests
import random
import string
import time

# Define the API endpoint
url = "http://localhost:8080/api/message"

# Define the symbols to choose from
symbols = [
    "AAPL", "GOOG", "MSFT", "AMZN", "FB",
    "TSLA", "NVDA", "INTC", "AMD", "NFLX",
    "PYPL", "ADBE", "CSCO", "CMCSA", "PEP",
    "COST", "AMGN", "TXN", "AVGO", "QCOM",
    "SBUX", "GILD", "BKNG", "MDLZ", "CHTR",
    "REGN", "FISV", "ADP", "TMUS", "INTU",
    "ATVI", "ISRG", "CSX", "VRTX", "BIIB",
    "ILMN", "ADI", "MELI", "BIDU", "JD",
    "NXPI", "LRCX", "ADI", "KHC", "EXC",
    "EA", "CTSH", "WBA", "MAR", "ORLY",
    "KLAC", "WDC", "CDNS", "MNST", "CTAS",
    "XEL", "VRSK", "PCAR", "FAST", "SIRI",
    "CDW", "ANSS", "SWKS", "MXIM", "NTES",
    "ALGN", "SNPS", "ALXN", "ULTA", "IDXX",
    "MCHP", "TTWO", "XLNX", "CERN", "VRSN",
    "WYNN", "LULU", "PAYX", "BMRN", "TCOM",
    "CHKP", "ASML", "CPRT", "NTAP", "CTRP",
    "EXPE", "NTNX", "MTCH", "ZM", "OKTA"
]

# Define a function to generate a random symbol
def random_symbol():
    return random.choice(symbols)

# Define a function to generate a random value
def random_value():
    return random.randint(0, 1000)

# Define a function to generate a random JSON payload
def random_payload():
    return {
        "symbol": random_symbol(),
        "val": random_value()
    }

# Define a function to generate a random string
def random_string(length):
    return ''.join(random.choice(string.ascii_letters) for i in range(length))

# Define a function to generate a random JSON payload with a random symbol
def random_payload_with_random_symbol():
    return {
        "symbol": random_string(3),
        "val": random_value()
    }

# Define the number of requests to send
num_requests = 1000

# Define the sleep time between requests (in seconds)
sleep_time = 0.01

# Send the requests
for i in range(num_requests):
    # Generate a random JSON payload
    payload = random_payload_with_random_symbol()

    # Send the POST request
    response = requests.post(url, json=payload)

    # Print the response status code
    print(f"Request {i+1}: {response.status_code}")

    # Sleep for a short time to avoid overwhelming the API
    time.sleep(sleep_time)
