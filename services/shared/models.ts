interface Product extends EntityBase {
    category: Category[],
    dimensions: { width: number, height: number, depth: number, unit: 'cm' }
    weight: { quantity: number, unit: 'kg' },
    availability: {
        // - Each product item shares the product UUID and we only tracing ammount and location 
        // - (like in the grocery store, each product (battle of coke) item has the same barcode)
        location: {
            uuid: string,
            type: 'shop' | 'warehouse', // enum
        },
        stataus: 'ordered' | 'transport' | 'arrived' | 'processing' | 'available' | 'sold' // enum, enums are in time order
        quantity: number,
        timeOfInbound: Timestamp,
    }[] // there could be multiple quantity on multiple location
    status: 'AVAILABLE' | 'OUT_OF_STOCK' | 'ORDERABLE' | 'NOT_ORDERABLE'
}

interface Category extends EntityBase {
    subs: Category[],
}

interface EntityBase {
    uuid: string,
    timestamp: Timestamp,
    name: string,
    tags: string[],
    metadata: { 
        // metadata could be also used to store custom field from data migration
        [key: string]: MetaData | {} | [] | string 
    },
    description: string,
}

type Timestamp = string
type MetaData = {}
