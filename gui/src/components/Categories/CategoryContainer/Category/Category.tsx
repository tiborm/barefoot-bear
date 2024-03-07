import { Category } from "@core/models/category/category";
import { useEffect, useState } from "react";

interface CategoryContainerProps {
    category: Category;
}

export default function CategoryItem({ category }: CategoryContainerProps) {

    const [randomNumberOfProducts, setRandomNumberOfProducts] = useState<number>(0);

    useEffect(() => {
        setRandomNumberOfProducts(Math.floor(Math.random() * 200) + 1 + 120);
    });

    return (
        <li key={category.id} className="flex justify-between gap-x-6 py-5">
            <div className="flex min-w-0 gap-x-4">
                <img
                    className="h-12 w-12 flex-none rounded-full bg-gray-50 border-solid border-2 border-amber-500"
                    src={category.im}
                    alt=""
                />
                <div className="min-w-0 flex-auto">
                    <p className="text-sm font-semibold leading-6 text-amber-500">
                        {category.name}
                    </p>
                    <p className="mt-1 text-xs leading-5 text-gray-500 max-w-md">
                        <span className={"font-semibold"}>Subcategories:</span>{" "}
                        {category.subs.map((subcategory, index) => {
                            return subcategory.name + ((index < category.subs.length - 1) ? ", " : "");
                        })}
                    </p>
                </div>
            </div>
            <div className="hidden shrink-0 sm:flex sm:flex-col sm:items-end">
                {category.url && (
                    <p className="text-sm leading-6 text-gray-900">
                        <a href={category.url} target="_blank">
                            Go to page
                        </a>
                    </p>
                )}
                <p className="mt-1 text-xs leading-5 text-gray-500">
                    Number of products:&nbsp;
                    <span className={"font-bold text-blue-700"}>
                        {randomNumberOfProducts}
                    </span>
                </p>
            </div>
        </li>
    );
}
