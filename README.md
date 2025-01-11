
# go-reloaded

![Go Text Processor](https://golang.org/doc/gopher/frontpage.png)

## Features

### 1. **Case Transformation**
- **`(up, n)`**: Converts the last `n` words to uppercase.
- **`(low, n)`**: Converts the last `n` words to lowercase.
- **`(cap, n)`**: Capitalizes the first letter of the last `n` words.
- Example:
  - Input: `hello world (up, 1)`
  - Output: `hello WORLD`

### 2. **Binary and Hexadecimal to Decimal Conversion**
- Converts numbers in **binary** or **hexadecimal** format to **decimal**.
- Example:
  - Input: `101 (bin)`
  - Output: `5`
  - Input: `1A (hex)`
  - Output: `26`

### 3. **Punctuation Fixes**
- Removes unnecessary spaces around punctuation marks.
- Ensures proper spacing after punctuation.
- Example:
  - Input: `Hello ,world !`
  - Output: `Hello, world!`

### 4. **Quotation Fixes**
- Adjusts spaces around single and double quotes for proper formatting.
- Example:
  - Input: ` "hello" world  `
  - Output: `"hello" world`

### 5. **Article Adjustments ("a" vs. "an")**
- Ensures proper use of "a" and "an" based on the following word.
- Handles silent "h" words like "hour," "honest," etc.
- Example:
  - Input: `a hour`
  - Output: `an hour`

### 6. **Bracket Spacing Fixes**
- Ensures proper spacing around text within parentheses.
- Example:
  - Input: `(hello world)`
  - Output: `(hello world)`

---

## Usage

### Prerequisites
- Go version 1.18 or later.

### Installation
Clone the repository:
```bash
https://01.tomorrow-school.ai/git/skulesho/go-reloaded.git
cd go-reloaded
```

### Running the Program
To run the program, use the following command:
```bash
go run . <input_file> <output_file>
```

#### Example:
Input file (`input.txt`):
```
Hello world! this is a test (up, 1)
```
Run the command:
```bash
go run . input.txt output.txt
```
Output file (`output.txt`):
```
Hello world! this is a TEST
```

---

## How It Works

### Main Workflow
1. **Input Reading**: The program reads the content of the input file.
2. **Text Processing**: The `fixText` function orchestrates multiple transformations on the text:
   - Case modifications
   - Binary/hexadecimal conversions
   - Spacing and punctuation fixes
   - Grammar adjustments
3. **Output Writing**: The modified text is written to the output file.

### Modular Design
The program is modular, with dedicated functions for each type of text transformation:
- `textModifyCase`: Handles case transformations.
- `hexAndBinToDecimal`: Converts binary/hexadecimal to decimal.
- `fixPunctuations`: Fixes punctuation spacing.
- `fixDoubleQuotes` / `fixSingleQuotes`: Adjusts quotes.
- `fixAtoAn`: Ensures proper use of "a" and "an".

---

## Example Output

### Input File
```
this is an test of (up, 2) feature. 101 (bin) and 1A (hex) (low, 1).
```

### Output File
```
this is a TEST OF feature. 5 and 26.
```

---


![Go Lang Logo](https://blog.golang.org/gopher/gopher.png)
