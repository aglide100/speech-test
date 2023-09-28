import os

TOKEN = os.getenv("TOKEN")

options = [
    ('grpc.keepalive_time_ms', 900000),
    ('grpc.keepalive_permit_without_calls', True)
]


