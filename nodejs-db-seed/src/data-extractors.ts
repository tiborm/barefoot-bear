export const extractProducts = (productsByCategoryJSONs: { results: []; }[]) => {
    const extractedProducts: {}[] = [];
    productsByCategoryJSONs.forEach(products => {
        products.results.forEach((result: { items: []; }) => {
            result.items.forEach((productWithMetaData: { product: {}; }) => {
                // there is an item called planner which is not a product
                if (productWithMetaData.product) {
                    extractedProducts.push(productWithMetaData.product);
                }
            });
        });
    });

    return extractedProducts;
};

export const extractInventory = (inventoryByProductJSON: { data: {}[]; }[]) => {
    return inventoryByProductJSON.map(inventory => inventory.data).flat();
};
