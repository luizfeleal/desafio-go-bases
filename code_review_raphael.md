# Code Review Go-Bases Raphael

# 1. Utilize a estrutura de coleta direta no csv em teste

Assim é possível validar o funcionamento de leitura dos dados e deixar mais dinâmico a alteração

```go
csvContent := `1,Alice,alice@example.com,Brazil,08:00,100.0
2,Bob,bob@example.com,USA,14:00,200.0
3,Carol,carol@example.com,Brazil,20:00,150.0
4,David,david@example.com,Brazil,05:30,120.0
```
