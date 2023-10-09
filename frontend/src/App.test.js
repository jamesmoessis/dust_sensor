import { render, screen } from '@testing-library/react';
import App from './App';

test('renders threshold title', () => {
  render(<App />);
  const thresTitle = screen.getByText(/Threshold/i);
  expect(thresTitle).toBeInTheDocument();
});
