Base URL: http://localhost:9000/

Insert Rule Template
Endpoint:
POST /insertRuleTemplate

Deskripsi:
Menambahkan template rule

Request Body:

Contoh Request:
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
    "condition": "Ini adalah deskripsi kondisi",
    "action": "Ini adalah deskripsi tindakan"
  }
}


Respon Sukses
Kode status: 201 Created
Contoh Body Respon
{
    "message": "rule set inserted"
}

Respon jika ruleset sudah ada:
{
    "message": "rule set already exist"
}


Insert Rules to Rule Set
Endpoint:
PATCH /insertRuletoRuleSet

Deskripsi:
Memasukkan rule baru ke dalam Rule Set

Parameter Query:
‘ruleSetName’: Nama dari Rule Set yang akan dimasukkan Rule

Contoh Request
[
  {
    "conditions": {
      "$hargalebih": 15
    },
    "action": 10
  }
]


Respon Sukses
Kode Status: 200 OK
Body Respon
{
“message”: “1 new rules has been inserted”
}

Executed Rule Set into Action
Endpoint
POST /execInput

Deskripsi:
Mengolah input dari user dengan rule set.

Parameter Query:
‘ruleSetName’: Nama dari Rule Set yang akan dimasukkan Rule

Payload:
{
    "nama": "suki",
    "harga": 15,
    "jualan": 20
}

Response:
200 OK

{
    "message": 30
}


Update Rule Set
Endpoint
PUT /updateRuleSet

Deskripsi
Meng-update rule-rule yang ada di dalamRule Set

Parameter Query:
‘ruleSetName’: Nama dari Rule Set yang akan dimasukkan Rule

Contoh Request:
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

Respon Sukses
Kode Status: 200 OK
Contoh Body Respon:
{
    "message": RuleSet updated successfully"
}

List All Rule Set
Endpoint
GET /fetchRules

Deskripsi
Mengambil semua rule set yang tersedia

Respon Sukses
Kode Status: 200 OK
Contoh Body Respon ketika rule set tersedia
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





Contoh Body Respon ketika rule set kosong
{
    "message": "rule set list empty"
}

Delete Rule Set from DB
Endpoint
DELETE /deleteRuleSet

Deskripsi:
Menghapus rule set dari mongodb


Parameter Query:
‘ruleSetName’: Nama dari Rule Set yang akan dimasukkan Rule

Response:
200 OK
{
    "message": "rule set RuleSet1 has been deleted"
}

