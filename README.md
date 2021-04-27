# prog-tester
## Usage
- tester \<command> \<?value> \<optional args>

## Available commands
- **run**: runs the test on given executable `tester run <executable>`
- **gen**: generates the required file structure, number represents the count of input entries, it's argument is optional `tester gen <?number>`
#
Planned commands:
- **diff**: takes in a testcase (it's name or id `tester diff 4`, `tester diff testcase4`) and prints out unified diff of the output and expected output or instead of a new command just print the diff to the output txt?
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
- Every element/tag inside a `<testcase>` is omittable, it means that an empty testscase is totally valid, so it will run a testcase with no arguments, empty input, no output is expected and the exitcode should be zero, otherwise the testcase will fail
- `<input>` and `<expect>` have attribute named **newline**, which is useful when you need enters at the end of you input (or expect), see in example
- Make sure your inputs and expects (only if they are on multiple lines, see last testcase in the example) are indented exactly as you want it!!!

## Example test_schema.xml
```xml
<?xml version="1.0" encoding="UTF-8"?>
<tester>
  <testcase>
    <args>-c -w</args>
    <exitcode>1</exitcode>
  </testcase>

  <testcase>
    <args>-l -u</args>
    <exitcode>3</exitcode>
  </testcase>

  <testcase>
    <args>-a</args>
    <input newline="2">jednoVelkeSpojeneSlovo</input>
    <expect newline="1">jednoVelkeSpojeneSlovo</expect>
  </testcase>

  <testcase>
    <args>-a</args>
    <input>slovo        ? + 1 aaa AAA</input>
    <expect>slovo              aaa AAA</expect>
  </testcase>

  <testcase>
    <args>-a</args>
    <input>   C84ha  .o+t ?i c    W[o ]r .D)</input>
    <expect>   C  ha   o t  i c    W o  r  D </expect>
  </testcase>

  <testcase>
    <args>-c</args>
    <input newline="2">   *-+ A b e c 8 A llllvs  [e][t]ko{}JeNa   t.lac89e741ne +V+  L??A?V ??O</input>
    <expect newline="1">AbecAllllvsetkoJeNatlaceneVLAVO</expect>
  </testcase>

  <testcase>
    <args>-c</args>
    <input newline="2">   *-+ A b e c 8 A llllvs  [e][t]ko{}JeNa   t.lac89e741ne +V+  L??A?V ??O</input>
    <expect newline="1">AbecAllllvsetkoJeNatlaceneVLAVO</expect>
  </testcase>

  <testcase>
    <args>-c</args>
    <input>The ladies of Longbourn soon waited on those of Netherfield  The
visit was soon returned in due form  Miss Bennets pleasing
manners grew on the goodwill of Mrs  Hurst and Miss Bingley  and
though the mother was found to be intolerable  and the younger
sisters not worth speaking to  a wish of being better acquainted
with  them  was expressed towards the two eldest  By Jane  this
attention was received with the greatest pleasure  but Elizabeth
still saw superciliousness in their treatment of everybody 
hardly excepting even her sister  and could not like them  though
their kindness to Jane  such as it was  had a value as arising in</input>
    <expect>TheladiesofLongbournsoonwaitedonthoseofNetherfieldThe
visitwassoonreturnedindueformMissBennetspleasing
mannersgrewonthegoodwillofMrsHurstandMissBingleyand
thoughthemotherwasfoundtobeintolerableandtheyounger
sistersnotworthspeakingtoawishofbeingbetteracquainted
withthemwasexpressedtowardsthetwoeldestByJanethis
attentionwasreceivedwiththegreatestpleasurebutElizabeth
stillsawsuperciliousnessintheirtreatmentofeverybody
hardlyexceptingevenhersisterandcouldnotlikethemthough
theirkindnesstoJanesuchasitwashadavalueasarisingin</expect>
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
  EXIT_CODE: got: 1, expected: 0
  PROGRAM_EXITED_WITH: segmentation fault (core dump)
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
  EXIT_CODE: got: 132543 expected: 0
  PROGRAM_EXITED_WITH: exit code 132543 
TestCase6..........OK
TestCase7..........OK
TestCase8..........OK
TestCase9..........OK
==================================================
Ran 9 tests (passed: 5, failed: 4)
TEST FAILED
```
### Output txt file
```
OUTPUT:
Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Sed eu nisi at nibh semper cursus ac et neque.
----------------------------------------------------------
EXPECTED:
Why would you think the doggy is little?
the doggy is the most big canis familiaris of all.
Are you upset by how monstrous it is?
Does it tear you apart to see the doggy so large-scale?
I cannot help but stop and look at the mad tabby.
```