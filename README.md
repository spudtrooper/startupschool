# startupschool

Tool to scrape startupschool.org founder matching (because I'd prefer to read them in batch & offline)

## Usage

1. Create a file `.credentials.json` with the following:

    ```json
    {
        "startupschoolUsername": "<email>",
        "startupschoolPassword": "<password>"
    }
    ```

2. Collect some candidates:

    ```bash
    ./scripts/main.sh [--limit <N>]
    ```

3. Create a report

    ```bash
    ./scripts/report.sh
    ```

4. Check out the report

    ```bash
    open data/html/index.html
    ```