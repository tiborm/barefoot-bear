import React from 'react';
import type { Preview } from "@storybook/react";

import '../styles/globals.css';

import { StoryContext, Story } from '@storybook/react';



const preview: Preview = {
  parameters: {
    decorators: [
      (Story: React.FC, context: StoryContext) => (
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
          <Story {...context.args} />
        </div>
      ),
    ],
    actions: { argTypesRegex: "^on[A-Z].*" },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
};

export default preview;

