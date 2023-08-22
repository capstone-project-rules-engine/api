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

**Parameter Query**: ‘ruleSetName’: Nama dari Rule Set yang akan dimasukkan Rule

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
    “message”: “1 new rules has been inserted”
}
```
