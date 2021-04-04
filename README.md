# prog-tester
## Usage
- tester \<command> \<?value> \<optional args>

## Available commands
- run -> runs the test on given executable `tester run <executable>`
- gen -> generates the required file structure, number represents the count of input entries, it's argument is optional `tester gen <?number>`

## Available flags
- --save-output, -s -> saves the output os the test case to a txt
- --name, -n        -> specify the name of the input xml file, by default it is test_schema.xml

## Required folder structure
```
WORK_DIR
├── test_schema.xml
├── your_executable
└── tester
```

## Input file format
Test cases are specified in xml file. Use this format to add test cases entries.
- NOTE: If exit code not given, the tester assumes it should be 0.
- NOTE: Make sure your inputs and expects (only if they are on multiple lines) are indented exactly as you want !!!
```xml
<?xml version="1.0" encoding="UTF-8"?>
<tester>
  <testcase>
    <args>-r priatel nepriatel</args>
    <input>The ladies of Longbourn soon waited on those of Netherfield  The
visit was soon returned in due form  Miss Bennets pleasing
manners grew on the goodwill of Mrs  Hurst and Miss Bingley  and
though the mother was found to be intolerable  and the younger
sisters not worth speaking to  a wish of being better acquainted
with  them  was expressed towards the two eldest  By Jane  this
attention was received with the greatest pleasure  but Elizabeth
still saw superciliousness in their treatment of everybody
hardly excepting even her sister  and could not like them  though
their kindness to Jane  such as it was  had a value as arising in
all probability from the influence of their brothers admiration
It was generally evident whenever they met  that he  did  admire
her and to  her  it was equally evident that Jane was yielding to
the preference which she had begun to entertain for him from the
first  and was in a way to be very much in love  but she
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