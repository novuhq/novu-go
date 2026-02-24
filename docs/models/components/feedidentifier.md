# FeedIdentifier

Optional feed identifier or array of feed identifiers


## Supported Types

### 

```go
feedIdentifier := components.CreateFeedIdentifierStr(string{/* values here */})
```

### 

```go
feedIdentifier := components.CreateFeedIdentifierArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch feedIdentifier.Type {
	case components.FeedIdentifierTypeStr:
		// feedIdentifier.Str is populated
	case components.FeedIdentifierTypeArrayOfStr:
		// feedIdentifier.ArrayOfStr is populated
}
```
