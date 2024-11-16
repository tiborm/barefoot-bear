package fetch

var SearchJSONTemplate = []byte(`{
    "searchParameters": {
        "input": "59308",
        "type": "CATEGORY"
    },
    "isUserLoggedIn": false,
    "optimizely": {
        "listing_1985_mattress_guide": null,
        "listing_fe_null_test_12122023": null,
        "listing_1870_pagination_for_product_grid": null
    },
    "components": [
        {
            "component": "PRIMARY_AREA",
            "columns": 4,
            "types": {
                "main": "PRODUCT",
                "breakouts": [
                    "PLANNER",
                    "LOGIN_REMINDER"
                ]
            },
            "filterConfig": {
                "max-num-filters": 7
            },
            "sort": "RELEVANCE",
            "window": {
                "size": 24,
                "offset": 0
            },
            "forceFilterCalculation": true
        }
    ]
}`)
