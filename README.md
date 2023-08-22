# API Documentation

## Base URL: `http://localhost:9000/`

### Insert Rule Template

**Endpoint**: `POST /insertRuleTemplate`

**Description**: Adds a new rule template.

**Request Body**:

**Example Request**:

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
