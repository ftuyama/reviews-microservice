db.createCollection("reviews");

function get_results(result) {
    print(tojson(result));
}

function insert_reviews(object) {
    print(db.reviews.insert(object));
}

insert_reviews({
    "customerId": "57a98d98e4b00679b4a830b2",
    "itemId": "837ab141-399e-4c1f-9abc-bace40296bac",
    "createdAt": new Date("2024-04-24T00:09:22.198Z"),
    "updatedAt": new Date("2024-04-24T00:09:22.198Z"),
    "rating": 5,
    "comment": "Excellent product!"
});

print("_______REVIEW DATA_______");
db.reviews.find().forEach(get_results);
print("______END REVIEW DATA_____");
