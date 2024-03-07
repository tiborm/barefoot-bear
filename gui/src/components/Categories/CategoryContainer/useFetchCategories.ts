import { Category } from "@core/models/category/category";
import { useEffect, useState } from "react";

const useFetchCategories = () => {
    const [categories, setCategories] = useState<Category[]>([]);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        //const fetchCategories = async () => {
        const fetchCategories = () => {
            // const response = await fetch("http://localhost:3001/categories");
            // NOTE embeedding the data directly into the component for now
            const data: Category[] = require("../../../../../data-transplant/json-cache/categories.json");
            setCategories(data);
            setLoading(false);
        }
        fetchCategories();
    }, []);

    return { categories, loading };
}

export default useFetchCategories;
