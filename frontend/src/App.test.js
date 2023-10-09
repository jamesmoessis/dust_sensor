import { render, screen } from '@testing-library/react';
import App from './App';

test('renders button title', () => {
  render(<App />);
  const buttonTitle = screen.getByText(/POWER/i);
  expect(buttonTitle).toBeInTheDocument();
});
