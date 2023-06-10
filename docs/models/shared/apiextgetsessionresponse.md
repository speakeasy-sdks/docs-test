# ApiextGetSessionResponse

OK


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `CreatedAt`                                                                              | **string*                                                                                | :heavy_minus_sign:                                                                       | When this session was created.                                                           | 2022-01-11 22:32:45.601486+00                                                            |
| `CreatedBy`                                                                              | **string*                                                                                | :heavy_minus_sign:                                                                       | ID of the user that created this session.                                                | usr20220103zlufhym                                                                       |
| `ID`                                                                                     | **string*                                                                                | :heavy_minus_sign:                                                                       | Unique ID of this session.                                                               | ses20220120za1pskd                                                                       |
| `IsPrivate`                                                                              | **bool*                                                                                  | :heavy_minus_sign:                                                                       | Whether or not the session is private.                                                   | true                                                                                     |
| `Name`                                                                                   | **string*                                                                                | :heavy_minus_sign:                                                                       | Name of this session.                                                                    | MySession                                                                                |
| `ParamValues`                                                                            | map[string]*string*                                                                      | :heavy_minus_sign:                                                                       | Mapping of parameter slug to value used in this session's execution.                     |                                                                                          |
| `Params`                                                                                 | [][ApiextParameter](../../models/shared/apiextparameter.md)                              | :heavy_minus_sign:                                                                       | Schema for the set of values users can provide when executing this session.              |                                                                                          |
| `Permissions`                                                                            | [][ApiextPermission](../../models/shared/apiextpermission.md)                            | :heavy_minus_sign:                                                                       | Explicit permissions of this session if it is private.                                   |                                                                                          |
| `RunbookID`                                                                              | **string*                                                                                | :heavy_minus_sign:                                                                       | ID of the runbook this session was spawned from if triggered from a runbook.             | rbk20220120z15kl79                                                                       |
| `Status`                                                                                 | [*ApiextGetSessionResponseStatus](../../models/shared/apiextgetsessionresponsestatus.md) | :heavy_minus_sign:                                                                       | Status of this session.                                                                  | Succeeded                                                                                |
| `TeamID`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | ID of the team that owns this session.                                                   | tea20220103zvy4auu                                                                       |
| `UpdatedAt`                                                                              | **string*                                                                                | :heavy_minus_sign:                                                                       | When this session was updated.                                                           | 2022-01-11 22:35:45.238512+00                                                            |
| `UpdatedBy`                                                                              | **string*                                                                                | :heavy_minus_sign:                                                                       | ID of the user who updated this session.                                                 | ses20220120za1pskd                                                                       |