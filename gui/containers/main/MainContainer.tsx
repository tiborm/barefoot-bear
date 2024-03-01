import React from "react";
import categories from "../../../json-cache/categories.json";

import styles from "./styles.module.scss";

export function MainContainer() {
    return (
        <div className={styles.mainContainer}>
            <ul>
                {
                    categories.map((category) => {
                        return <li key={category.id}>{category.name}</li>;
                    })
                }
            </ul>
        </div>
    );
}
