import { render } from "@testing-library/react";
import { expect, describe, it } from "vitest";
import { MainContainer } from "..";
import sinon from "sinon"
import * as CategoryContainer from "../Categories/CategoryContainer/CategoryContainer";

sinon.stub(CategoryContainer, 'default').returns(<div data-testid="category-container"></div>);

describe("MainContainer", () => {
    it("Should render CategoryContainer", () => {
        const { getByTestId } = render(<MainContainer />);
    
        expect(getByTestId("category-container")).toBeTruthy();
    });

});
