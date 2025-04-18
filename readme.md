# go-generator

🚀 **go-generator** is a lightweight CLI tool written in Go for generating boilerplate files (models, services, handlers, etc.) based on embedded templates and a simple YAML configuration file.

---

## 📦 Installation

Install the generator globally with:

```bash
go install github.com/Albert-Przybyla/go-generator/cmd/gen-struct@latest
```

## ⚙️ Configuration

To use the generator, create a .conf.yaml file in your project root. This file defines where different types of generated files should be placed.

Example **.conf.yaml**

```yaml
paths:
  models: internal/models
  repositories: internal/repositories
  services: internal/services
  dtos: internal/dtos
  mappers: internal/mappers
  handlers: internal/handlers
```

Configuration Structure

| key          | Description                                              | Example Path          |
| ------------ | -------------------------------------------------------- | --------------------- |
| models       | Directory for generated model files                      | internal/models       |
| repositories | Directory for generated repository files                 | internal/repositories |
| services     | Directory for generated service files                    | internal/services     |
| dtos         | Directory for generated Data Transfer Object (DTO) files | internal/dtos         |
| mappers      | Directory for generated mapping logic files              | internal/mappers      |
| handlers     | Directory for generated HTTP handler files               | internal/handlers     |

## 🚀 Usage

Navigate to your project root (where .conf.yaml is located) and run:

```bash
gen-struct "User Profile"
```

Optionally, you can use the -simple flag for simpler template output:

```bash
gen-struct -simple "User"
```

This will generate

## 🧪 Example Output

Running:

```bash
gen-struct "Invoice Item"
```

Will generate:

- internal/models/invoice_item_model.go
- internal/repositories/invoice_item_repository.go
- internal/services/invoice_item_service.go
- internal/dtos/invoice_item_dto.go
- internal/mappers/invoice_item_mapper.go
- internal/handlers/invoice_item_handler.go

## 📜 License

MIT © Albert Przybyła
