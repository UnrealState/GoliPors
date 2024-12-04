POST /api/surveys
Content-Type: application/json

{
  "title": "Customer Feedback Survey",
  "creation_time": "2024-12-01T00:00:00Z",
  "start_time": "2024-12-05T00:00:00Z",
  "end_time": "2024-12-31T23:59:59Z",
  "random_order": true,
  "allow_return": false,
  "num_participation_attempts": 3,
  "response_time": 1800,
  "anonymity_level": "visible_to_creator_and_admins",
  "demographic_restrictions": "{\"age\": \"18-65\", \"location\": \"USA\"}",
  "response_modification": true
}
