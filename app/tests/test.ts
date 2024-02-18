import { expect, test } from '@playwright/test';

test('index page has stredono text', async ({ page }) => {
	await page.goto('/');
	await expect(page.getByText("Stredono")).toBeVisible();
});
