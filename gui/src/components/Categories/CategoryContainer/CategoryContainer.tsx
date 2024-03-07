import CategoryItem from "./Category/Category";
import useFetchCategories from "./useFetchCategories";

// import styles from "./styles.module.scss";

export default function CategoryContainer() {
    const { categories, loading } = useFetchCategories();

    return (
        <div
            data-testid="category-container"
            className={`flex min-h-screen flex-col items-center justify-between p-24`}
        >
            <ul role="list" className="divide-y divide-gray-100">
                {categories.map((category) => {
                    return (
                        <CategoryItem key={category.id} category={category} />
                    );
                })}
            </ul>
        </div>
    );
}
