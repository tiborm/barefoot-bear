import type { Meta, StoryObj } from "@storybook/react";

import MainContainer from "./MainContainer";

const meta: Meta<typeof MainContainer> = {
    title: "Example/MainContainer",
    component: MainContainer,
    // ...
};
export default meta;

type Story = StoryObj<typeof MainContainer>;

export const Primary: Story = {
    // args: {
    //     primary: true,
    //     label: "Click",
    //     background: "red",
    // },
};

export const Warning: Story = {
    // args: {
    //     primary: true,
    //     label: "Delete now",
    //     backgroundColor: "red",
    // },
};
