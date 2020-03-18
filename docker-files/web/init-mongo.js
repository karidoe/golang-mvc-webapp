db.createUser(
    {
        "user": "anuchit",
        "pwd": "$$jak$$",
        "roles": [
            {
                "role": "readWrite",
                "db": "test"
            }
        ]
    }
)