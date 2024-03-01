import { expect, test } from "vitest";
import { render, screen } from "@testing-library/react";
import { MainContainer } from "..";

test("MainContainer tests", () => {
    render(<MainContainer />);

    expect(true).toBe(true);
});
