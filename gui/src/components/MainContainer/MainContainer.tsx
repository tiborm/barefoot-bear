import React from "react";

import styles from "./styles.module.scss";
import { Category } from "@core/models/category/category";

// NOTE embeedding the data directly into the component for now
const categories: Category[] = require("../../../../data-transplant/json-cache/categories.json");

export function MainContainer() {
    return (
        <div className={styles.mainContainer}>
            <ul role="list" className="divide-y divide-gray-100">
                {categories.map((category) => {
                    return (
                        <li
                            key={category.id}
                            className="flex justify-between gap-x-6 py-5"
                        >
                            <div className="flex min-w-0 gap-x-4">
                                <img
                                    className="h-12 w-12 flex-none rounded-full bg-gray-50"
                                    src={category.im}
                                    alt=""
                                />
                                <div className="min-w-0 flex-auto">
                                    <p className="text-sm font-semibold leading-6 text-gray-900">
                                        {category.name}
                                    </p>
                                    <p className="mt-1 truncate text-xs leading-5 text-gray-500">
                                        leslie.alexander@example.com
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
                                    Last seen{" "}
                                    <time dateTime="2023-01-23T13:23Z">
                                        3h ago
                                    </time>
                                </p>
                            </div>
                        </li>
                    );
                })}
            </ul>
        </div>
    );
}
