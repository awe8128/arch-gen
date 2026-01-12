# Project Generator (YAML → DDD / Code Skeleton)

A CLI tool that generates backend project structure and boilerplate code from a single YAML definition.

This project is designed for teams who want:
consistent folder structure (ex: DDD)
consistent entity/repository shapes
one “source of truth” (YAML) for domain models + DB models + repository signatures

## What it generates
From config.yaml (or your chosen file), the generator can create:
Project root based on pj.name
System structure based on pj.sys (ex: ddd)
Domain folders for each entry under domains
Entity structs with nullable handling
Repository interface/function templates based on repositories
(Optional) DB table model stubs based on db
The goal is not to generate your whole app, but to generate the boring repeatable scaffolding fast and consistently.

## Quick Start

Create a YAML file (example below) named config.yaml

Run the generator

# Example (adjust to your actual command)

```go run ./cmd/... --config ./config.yaml```

# Or if you build it:

```./your-cli generate --config ./config.yaml```

Check the output directory
A new folder will be created using pj.name (example: be/)
Inside it, folders/files will be generated depending on pj.sys and domains

| Key       | Type   | Description                                  |
| --------- | ------ | -------------------------------------------- |
| `pj`      | object | project settings (name, system type)         |
| `db`      | object | DB table-like definitions (simple)           |
| `domains` | object | domain definitions (entities + repositories) |

```
pj:
  name: be
  sys: ddd

db:
  user:
    id: uuid
    name: string
  lead:
    id: int64
    price: int64
    name: string

domains:
  User:
    type: struct
    properties:
      id:
        type: uint
        nullable: false
      name:
        type: string
        nullable: true
      age:
        type: int64
        nullable: false
    repositories:
      Create:
        type: func
        in:
          id:
            type: uint
            nullable: false
          name:
            type: string
            nullable: true
        out:
          name:
            type: string
            nullable: true
          id: 
            type: uuid.UUID
            nullable: false

  Lead:
    type: struct
    properties:
      id:
        type: uint
        nullable: true
```
