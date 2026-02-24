# Tenant

It is used to specify a tenant context during trigger event.
    Existing tenants will be updated with the provided details.


## Supported Types

### 

```go
tenant := components.CreateTenantStr(string{/* values here */})
```

### TenantPayloadDto

```go
tenant := components.CreateTenantTenantPayloadDto(components.TenantPayloadDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch tenant.Type {
	case components.TenantTypeStr:
		// tenant.Str is populated
	case components.TenantTypeTenantPayloadDto:
		// tenant.TenantPayloadDto is populated
}
```
