# prog-tester
## Usage
- tester \<command> \<?value> \<optional args>

## Available commands
- **run**: runs the test on given executable `tester run <executable>`
- **gen**: generates the required file structure, number represents the count of input entries, it's argument is optional `tester gen <?number>`

## Available flags
- **--discard-output**, **-d**: the produced output of executable won't be saved (By default every output produced from each testcase is saved to individual txt files)
- **--name**, **-n**: specify the name of the input xml file, by default it is test_schema.xml

## Required folder structure
```
WORK_DIR
├── test_schema.xml
├── your_executable
├── tester
└── outputs/ (This folder will be auto generated, if -d not provided)
```

## Input file format
Test cases are specified in xml file. Use this format to add test cases entries.
#### NOTES
- If exit code not given, the tester assumes it should be 0.
- Every element/tag inside a <testcase> is omittable, it means that an empty testscase is totally valid, so it will run a testcase with no arguments, empty input, no output is expected and the exitcode should be zero, otherwise the testcase will fail
- Make sure your inputs and expects (only if they are on multiple lines) are indented exactly as you want it!!!

## Example test_schema.xml
```xml
<?xml version="1.0" encoding="UTF-8"?>
<tester>
  <testcase>
    <args>-r priatel nepriatel</args>
    <input>The ladies of Longbourn soon waited on those of Netherfield  The
visit was soon returned in due form  Miss Bennets pleasing
manners grew on the goodwill of Mrs  Hurst and Miss Bingley  and
considered with pleasure that it was not likely to be discovered
by the world in general  since Jane united  with great strength
of feeling  a composure of temper and a uniform cheerfulness of
manner which would guard her from the suspicions of the
impertinent  She mentioned this to her friend Miss Lucas</input>
    <expect>Mrs  Hurst and Miss Bingley  and
though the mother was found to be intolerable  and the younger
sisters not worth speaking to  a wish of being better acquainted
with  them  was expressed towards the two eldest  By Jane  this
attention was received with the greatest pleasure  but Elizabeth
still saw superciliousness in their treatment of everybody</expect>
    <exitcode>1</exitcode>
  </testcase>

  <testcase>
    <args>-r "Godzilla vs. Kong" zadanie_4</args>
    <input>Velmi sa tesim na zadanie_4. Len ci sa take zadanie da vobec pripravit.</input>
    <expect>Len ci sa take zadanie da vobec pripravit.</expect>
    <exitcode>0</exitcode>
  </testcase>
</tester>
```
## Example output
### Terminal:
```
TestCase1..........OK
TestCase2..........FAIL
  args: lorem -a --help
  input: dolor sit amet
  ERROR: segmentation fault (core dump)
  EXIT_CODE: got: 1, expected: 0
TestCase3..........FAIL
  args: lorem -a --help
  input: dolor sit amet
  OUTPUT: not matching with expected
TestCase4..........FAIL
  args: lorem -a --help
  input: dolor sit amet
  EXIT_CODE: got: 0, expected: 2
TestCase5..........FAIL
  args: lorem -a --help
  input: dolor sit amet
  OUTPUT: not matching with expected
  EXIT_CODE: got: -132543 expected: 0
TestCase6..........OK
TestCase7..........OK
TestCase8..........OK
TestCase9..........OK
============================================
Ran 9 tests (passed: 5, failed: 4)
TEST FAILED
```
### Output txt file
```
OUTPUT:
Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Sed eu nisi at nibh semper cursus ac et neque.

EXPECTED:
Why would you think the doggy is little?
the doggy is the most big canis familiaris of all.
Are you upset by how monstrous it is?
Does it tear you apart to see the doggy so large-scale?
I cannot help but stop and look at the mad tabby.
```