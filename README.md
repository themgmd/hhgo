# HeadHunter SDK

## Installation

```sh
go get -u github.com/onemgvv/hhgo
```

## Usage

```go
otps := headhunter.Options{
    PublicId: "public-id-from-hh",
    PrivateKey: "private-key-from-hh",
}

hh, err := headhunter.New(opts)
if err != nil {
	// do something
}

dto := headhunter.SearchVacancyDto{
	Text: "Golang разработчик",
	Salary: 220000,
}

data := hh.GetVacancies(dto)
```