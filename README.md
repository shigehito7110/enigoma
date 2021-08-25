# Enigoma
## introduction
Enigoma is go package to hundle password(validate, convert to hash, compaire hash and password).

golang v1.16

## Usage
* use default validator func
```go
e, _ := enigoma.NewEnigoma(<password>)
if err := e.Validate(); err != nil {
  // error hundling
}

hash, err := e.CreateHash()
if err != nil {
  // error hundling
}

if err := e.Check(hash); err != nil {
  // error hundling
}
```
* use custom validator func
```go
type CustomValidator struct {}

func (c *CustomValidator) Validate(p string) bool {
  // validatation hundling
}

e, _ := enigoma.NewEnigoma(<password>, func() enigoma.vl.Validator {
  return &customValidator{
    // each fields
  }
})
```

## Default Validator

|  uppercase  |  downcase  |  numbers  |
| :---------: | :---------:| :--------:|
|      ◯      |     ◯      |     ◯     |

in 8 to 100 words

* example
```go
password // invalid
Password123 // correct
```
