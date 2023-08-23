# API Documentation

## Base URL: `http://localhost:9000/`

### Insert Rule Template

**Endpoint**: `POST /insertRuleTemplate`

**Deskripsi**: Menambahkan Template Rule.

**Contoh Request**:

```json
{
  "name": "RuleSet 1",
  "endpoint": "ruleset1",
  "bodies": [
    {
      "name": "harga",
      "type": "number"
    },
    {
      "name": "jualan",
      "type": "number"
    }
  ],
  "conditions": [
    {
      "label": "$hargalebih",
      "attribute": "harga",
      "operator": ">"
    },
    {
      "label": "$jualanlebih",
      "attribute": "jualan",
      "operator": ">"
    }
  ],
  "action": {
    "label": "$diskon",
    "attribute": "diskon",
    "type": "number"
  },
  "description": {
    "condition": "This is the condition description",
    "action": "This is the action description"
  }
}
```

**Respon Sukses**: 201 Created

**Contoh Body Respon**:

```json
{
  "message": "rule set inserted"
}
```

**Contoh Respon jika ruleset sudah ada**:

```json
{
  "message": "rule set already exist"
}
```

## Insert Rules to Rule Set

**Endpoint**: `PATCH /insertRuletoRuleSet`

**Deskripsi**: Memasukkan rule baru ke dalam Rule Set

**Parameter Query**: `ruleSetName`: 'endpoint' dari Rule Set yang akan dimasukkan Rule

**Contoh Request**:

```json
[
  {
    "conditions": {
      "$hargalebih": 15
    },
    "action": 10
  }
]
```

**Respon Sukses**: 200 OK

**Body Respon**:

```json
{
  "message": "1 rule set already exist"
}
```

## Executed Rule Set into Action

**Endpoint**: `POST /execInput`

**Deskripsi**: Mengolah input dari user dengan rule set.

**Parameter Query**: ‘ruleSetName’: 'endpoint' dari Rule Set yang akan dimasukkan Rule

**Contoh Request**:

```json
{
  "nama": "suki",
  "harga": 15,
  "jualan": 20
}
```

**Respon Sukses**: 200 OK

**Body Respon**:

```json
{
  "message": 30
}
```

## Update Rule Set

**Endpoint**: `PUT /updateRuleSet`

**Deskripsi**: Meng-update rule-rule yang ada di dalamRule Set

**Parameter Query**:`ruleSetName``: 'endpoint' dari Rule Set yang akan dimasukkan Rule

**Contoh Request**:

```json
{
  "name": "RuleSet 1",
  "endpoint": "ruleset1",
  "bodies": [
    {
      "name": "harga",
      "type": "number"
    },
    {
      "name": "jualan",
      "type": "number"
    }
  ],
  "conditions": [
    {
      "label": "$hargalebih",
      "attribute": "harga",
      "operator": ">"
    },
    {
      "label": "$jualanlebih",
      "attribute": "jualan",
      "operator": ">"
    }
  ],
  "action": {
    "label": "$diskon",
    "attribute": "diskon",
    "type": "number"
  },
  "rules": [
    {
      "conditions": {
        "$hargalebih": 12
      },
      "action": 25
    }
  ],
  "description": {
    "condition": "Deskripsi kondisi yang diupdate",
    "action": "Deskripsi tindakan yang diupdate"
  }
}
```

**Respon Sukses**: 200 OK
**Contoh Body Respon**:

```json
{
  "message": "RuleSet updated successfully"
}
```

**note**:
Jika mau update, karena ini menggunakan PUT request, maka harus di bawa seluruh strukturnya kecuali untuk field yang akan di ganti. Misal jika ingin mengganti 'action' fieldnya saja, maka masukan 'action' field yang sudah siap untuk di ganti namun untuk fields lain harus tetap ada dan sama

## List All Rule Set

**Endpoint**: `GET /fetchRules`

**Deskripsi**: Mengambil semua rule set yang tersedia

**Respon Sukses**: 200 OK
**Contoh Body Respon ketika rule set tersedia**

```json
{
  "message": "listing all rule sets",
  "details": [
    {
      "name": "RuleSet 1",
      "endpoint": "ruleset1",
      "bodies": [...],
      "conditions": [...],
      "action": {...},
      "rules": [...],
      "description": {...}
    }
  ]
}
```

**Contoh Body Respon ketika rule set kosong**

```json
{
  "message": "rule set list empty"
}
```

**Contoh body response ketika rule set ada**

```json
{
    "details": [...],
    "message": "x rule sets printed" // x adalah jumlah rule set
}

```

## Fetch specific rule set

**Endpoint**: `GET /fetchSpecificRuleSet`

**Deskripsi**: Mengambil satu rule set berdasarkan namanya

**Parameter Query**:
`ruleSetName`: 'endpoint' dari Rule Set yang akan dimasukan

**response**: 200 OK

```json
{
    "details":{...},
    "message": "printing rule set x" // di mana x adalah nama dari rule set
}

```

## Delete Rule Set from DB
**Endpoint**: `DELETE /deleteRuleSet`

**Deskripsi**: Menghapus rule set dari mongodb

**Parameter Query**:
`ruleSetName`: 'endpoint' dari Rule Set yang akan dimasukkan Rule

**Response**: 200 OK

```json
{
  "message": "rule set RuleSet1 has been deleted"
}
```