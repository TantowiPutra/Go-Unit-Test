# Go-Unit-Test
Repository To Learn Go Unit Test

Struktur Project Go: Controller, Service, Repository


Software Testing is Done to ensure application quality
Cost goes up as the pyramid goes up
<!-- Test Pyramid -->
.UI                 -> Buka website melalui browser, input form, dll, termasuk sistem backend (service & unit) 
...Service          -> Fokus ke 1 Aplikasi (Front-end + Back-end), bisa juga test langsung backend melalui API, keuntungannya tidak perlu running keseluruhan seperti test UI
                       cons: harus running database
......Unit          -> Check satu unit dari banyak unit pada 1 Aplikasi, misal Unit Test Controller, Unit Test Service, dan Repository


Unit is the smallest component

<!-- testing.T -->
Sebuah struct yang digunakan untuk unit test di Go-Lang

<!-- testing.M -->
Sebuah struct yang digunakan untuk mengatur life cycle testing

<!-- testing.B -->
Sebuah struct yang digunakan untuk benchmarking, mengetahui seberapa cepat dan efisien kode program dapat berjalan

<!-- Command To Run Unit Test -->
- go test                       -> Run unit test without showing corresponding functions
- go test -v                    -> Run unit test and show corresponding functions (v : verbose)
- go test -run <functionName>   -> Run unit test on spesific function

<!-- To Fail Unit test -->
- Fail()    -> Marks The Test as Failed, but Program still continue executing
- FailNow() -> Marks The Test as Failed, and stop the program immediately
- Error()   -> Marks The Test as Failed, make a report, but does not stop the program from executing
- Fatal()   -> Marks The Test as Failed, make a report, stop the program from executing

<!-- Unit testing using if-else is not efficient, especially when more data is involved, therefore we can utilize assertion-->
- In Go, Assertion library is not yet available, therefore it must be created manually
- TESTIFY is one of the most famous assertion library (https://github.com/stretchr/testify)
- To add the TESTIFY dependency, use type command "go get github.com/stretchr/testify"
- Testify has two packages, assert and require, assert calls (Fail()), and require calls (FailNow()) when the unit test fails.

<!-- Skip Unit Test -->
- testing package has method t.Skip() incase certain unit test must be skipped, for example when a unit test can only be run on a spesific operating system (MacOS/Windows)

<!-- Before and After Test (testing.M)-->
- Sometimes in a unit test, we're trying to do something before and after a unit test is executed
- If the code made before and after a certain unit test is identical, then it could lead to code inefficiency
- to solve this, we can use testing.M method to set the executing unit test
- To se the executing unit test, function named "TestMain" must be declared with the parameter testing.M
- Function TestMain is only executed ONCE per package

<!-- Subtest -->
- Golang supports the feature of creating unit test function IN ANOTHER unit test
- This feature rarely present / exist in another programming language

<!-- Table Test -->
- Using sub test, we can create dynamic unit test by preparing test data in a slice, then iterate through, this is known as table test

<!-- Mock -->
- Mock is an objec that has been programmed in a certain way, so that when it's called it will result to the data that has been programmed before
- Mock is a technique in unit testing, therefore we can make mock object from an object that is hard to test
- For example when we want to create a unit test, but the code program somehow oblige use to make a call to an API third party service, it will be difficult since we need to always call the third party service
  and the result could be of not what is desired.
- In this situation, it's best suited to use the Mock object.
- we can use TESTIFY to create mock object


