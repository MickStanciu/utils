# utils

## mapper 
### GenericMapper
GenericMapper will map from []T to []R by providing a transformation function

Example:

```go
transformer := func(data from) *to {
  return &to{
    FIELD1: data.FIELD1,
    FIELD2: fmt.Sprintf("%s - %d", data.Meta.FIELD1, data.Meta.FIELD2),
  }
}
res := mapper.GenericMapper(req, transformer)
```
