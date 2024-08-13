FuzzForge 🚀

FuzzForge is a high-speed Go-based directory and parameter fuzzer designed for security professionals and bug bounty hunters. It performs efficient directory brute-forcing and parameter testing with customizable wordlists and file extensions.
Features ✨

    High-Speed Fuzzing: Optimized for fast performance with concurrency support.
    Flexible Wordlists: Supports any user-specified wordlist for directory brute-forcing.
    Customizable Extensions: Allows specifying file extensions for directory testing.
    Output to File: Option to save results to an output file.
    User-Friendly: Simple command-line interface with clear options.

Installation 🛠️

    Clone the Repository:

    bash

git clone https://github.com/z3r0X0r/fuzzforge.git
cd fuzzforge

Build the Tool:
Make sure you have Go installed. Build the tool using:

bash

    go build -o fuzzforge

Usage 📜

bash

./fuzzforge -u <target-url> -w <wordlist-file> -x <extensions> -c <concurrency> -m <mode> -o <output-file>

Options ⚙️

    -u <target-url>: Target URL for fuzzing (e.g., http://example.com).
    -w <wordlist-file>: Path to the wordlist file (e.g., dirb_wordlist.txt).
    -x <extensions>: Comma-separated list of file extensions to test (e.g., .php,.html,.js,.txt).
    -c <concurrency>: Number of concurrent requests (e.g., 100).
    -m <mode>: Fuzzing mode: dir for directory fuzzing or param for parameter fuzzing.
    -o <output-file>: File to save results (e.g., output.txt).

Example Command 💻

bash

./fuzzforge -u http://example.com -w /path/of/the/wordlist

This command performs directory brute-forcing on http://example.com, using the wordlist dirb_wordlist.txt, testing file extensions .php, .html, .js, and .txt, with a concurrency of 100, and saving the results to output.txt.
Contributing 🤝

Contributions are welcome! Please open an issue or submit a pull request if you have improvements or bug fixes.
License 📜

This project is licensed under the MIT License. See the LICENSE file for details.
Contact 📧

For questions or support, please contact maxuzumaki888@gmail.com.
