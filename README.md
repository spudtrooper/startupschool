# startupschool

Tool to scrape startupschool.org founder matching (because I'd prefer to read them in batch)

## Usage

* Create a file .credentials.json with the following:

    ```json
    {
        "startupschoolUsername": "<email>",
        "startupschoolPassword": "<password>"
    }
    ```

* Collect some candidates:

    ```bash
    ./scripts/main.sh [--limit <N>]
    ```

* Create a report

    ```bash
    ./scripts/report.sh
    ```

* Check out the report

    ```bash
    open data/html/index.html
    ```