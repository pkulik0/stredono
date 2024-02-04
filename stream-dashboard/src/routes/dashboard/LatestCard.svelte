<script>
    import {Button, Card, Dropdown, DropdownItem, Listgroup} from "flowbite-svelte";
    import {DotsVerticalSolid, RedoOutline, RefreshOutline, ShieldSolid} from "flowbite-svelte-icons";

    let list = [
        {
            name: 'Neil Sims',
            email: 'email@windster.com',
            value: '$320',
            message: 'I hope this helps'
        },
        {
            name: 'Bonnie Green',
            email: 'email@windster.com',
            value: '$3467',
            message: 'Love what you are doing'
        },
        {
            name: 'Michael Gough',
            email: 'email@wieeendster.com',
            value: '$67',
            message: ''
        },
        {
            name: 'Evan Bochniak',
            email: 'ueou4@euuoaeueoa.com',
            value: '$100',
            message: 'Nothing beats a good cause'
        },
        {
            name: 'Mike Nichs',
            email: 'email@uuuu.com',
            value: '$222',
            message: 'aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 aha22 '
        }
    ];

    let menuButtonClass = "flex";
</script>

<Card padding="xl" size="xl" class="mt-4 flex-1">
    <div class="flex justify-between items-center mb-4">
        <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white">Latest Donations</h5>
        <a href="/donations" class="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500"> View all </a>
    </div>

    <Listgroup items={list} let:item class="border-0 dark:!bg-transparent">
        <div class="flex items-center space-x-4 rtl:space-x-reverse">
            <div class="flex-[0.25]">
                <p class="text-sm font-medium text-gray-900 truncate dark:text-white">
                    {item.name}
                </p>
                <p class="text-sm text-gray-500 truncate dark:text-gray-400">
                    {item.email}
                </p>
            </div>
            <div class="flex-[0.75]">
                {item.message}
            </div>
            <div class="inline-flex text-base font-semibold text-gray-900 dark:text-white">
                {item.value}
            </div>
            <Button size="sm" color="primary" outline class="text-xs">
                <DotsVerticalSolid class="dots-menu" />
            </Button>
            <Dropdown>
                <DropdownItem class="{menuButtonClass}">
                    <RefreshOutline/>
                    Refresh
                </DropdownItem>
                <DropdownItem class="{menuButtonClass}">
                    <ShieldSolid/>
                    Block
                </DropdownItem>
                <DropdownItem slot="footer" class="{menuButtonClass} text-red-500">
                    <RedoOutline/>
                    Refund
                </DropdownItem>
            </Dropdown>
        </div>
    </Listgroup>
</Card>