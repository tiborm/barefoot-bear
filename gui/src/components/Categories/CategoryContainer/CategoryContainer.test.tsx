import { render } from "@testing-library/react";
import { describe, it, expect, beforeEach, afterEach } from "vitest";
import CategoryContainer from "./CategoryContainer";
import { Category } from "@core/models/category/category";
import * as useFetchCategories from "./useFetchCategories";
import sinon from "sinon";
import * as CategoryItem from "./Category/Category";

describe("CategoryContainer", () => {
    let sandbox: sinon.SinonSandbox;

    beforeEach(() => {
        sandbox = sinon.createSandbox();
    });

    afterEach(() => {
        sandbox.restore();
    });

    it("Should render category items", () => {
        const testCategories: Category[] = [
            {
                id: "1",
                name: "Test Category 1",
                subs: [],
                url: "test-category-1",
                im: "Test Image 1",
            },
            {
                id: "2",
                name: "Test Category 2",
                subs: [],
                url: "test-category-2",
                im: "Test Image 2",
            },
        ];

        const subComponentStub = sinon
            .stub(CategoryItem, "default")
            .returns(<div data-testid="category-item"></div>);
        const useStub = sinon
            .stub(useFetchCategories, "default")
            .returns({ categories: testCategories, loading: false });

        const getByTestId = render(<CategoryContainer />);

        expect(subComponentStub.calledTwice).toBeTruthy();

        expect(getByTestId.getAllByTestId("category-item")).toHaveLength(2);

        subComponentStub.restore();
        useStub.restore();
    });
});
