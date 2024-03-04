<script lang="ts">
    import { saveAlert } from '$lib/alerts';
    import GifPicker from '$lib/comp/GifPicker.svelte';
    import SoundPicker from '$lib/comp/SoundPicker.svelte';
    import { Alert, Alignment, AnimationType, Currency, Event, EventType, Position, Speed } from '$lib/pb/stredono_pb';
    import { userStore } from '$lib/user';
    import { pbEnumToItems } from '$lib/util';
    import {
        Button,
        ButtonGroup, Checkbox,
        CloseButton,
        Drawer,
        Heading,
        Hr,
        Input,
        InputAddon,
        Label, Radio, RadioButton, Range,
        Select,
        Textarea
    } from 'flowbite-svelte';
    import {
        AlignCenterSolid,
        BarsFromLeftSolid,
        BarsSolid,
        BellActiveSolid,
        ImageOutline,
        PlusSolid,
        VolumeUpSolid
    } from 'flowbite-svelte-icons';
    import { onMount } from 'svelte';
    import { sineIn } from 'svelte/easing';
    import AlertViewer from './AlertViewer.svelte';

    export let hidden = true;
    export let eventType: EventType;

    let alert = new Alert();
    let startValue = "0";
    let endValue = "10";
    $: {
        alert.EventType = eventType;
        alert.Min = Number.parseFloat(startValue)
        alert.Max = Number.parseFloat(endValue)
        alert.TextColor = "#FFFFFF";
        alert.AccentColor = "#2381f8";
        alert.Message = "[user] donated [value] [currency]";
        alert.Animation = AnimationType.PULSE;
        alert.AnimationSpeed = Speed.MEDIUM;
        alert.Alignment = Alignment.JUSTIFY;
        alert.TextPosition = Position.BOTTOM;
    }

    let exampleEvent = new Event();
    $: {
        exampleEvent.Type = eventType;
        const amount = Math.round(alert.Min + (alert.Max - alert.Min) / 2);
        exampleEvent.Data = {
            "user": "John",
            "value": (isNaN(amount) ? 2137 : amount).toString(),
            "currency": currencySymbol,
            "message": "This is an example message. See which settings you like the most and adjust them to your needs."
        }
    }

    const addNew = async () => {
        if (!alert.GifUrl) {
            console.error("GIF is missing");
            return;
        }

        try {
            await saveAlert(alert);
            hidden = true;
        } catch (e) {
            console.error(e); // TODO: Show error to user
        }
    }

    $: animations = pbEnumToItems(AnimationType);

    let gifPickerOpen = false;
    let soundPickerOpen = false;

    let currency: Currency|undefined = undefined;
    let currencySymbol = "?";
    $: if(currency) {
        switch(currency) {
            case Currency.PLN:
                currencySymbol = "zł";
                break;
            default:
                currencySymbol = "?";
                break;
        }
    }

    onMount(() => {
        return userStore.subscribe((user) => {
            if (!user) {
                currency = undefined;
                return;
            }
            currency = user.Currency;
        });
    })

    let transitionParams = {
        x: 300,
        duration: 200,
        easing: sineIn,
    };
    let divClass = "overflow-y-auto z-50 p-4 bg-gray-100 dark:bg-gray-800";
</script>

<Drawer activateClickOutside={false} transitionType="fly" {transitionParams} bind:hidden {divClass} placement="right" width="w-96">
    <div class="flex items-center">
        <h5 id="drawer-label" class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400">
            <BellActiveSolid class="w-4 h-4 me-2.5" />
            New {JSON.parse(JSON.stringify(EventType))[eventType]} Alert
        </h5>
        <CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
    </div>

    <div class="w-full space-y-4">
        <Heading tag="h5">Trigger</Heading>

        <div class="space-x-4 flex flex-row">
            <Label>
                From
                <ButtonGroup>
                    <Input type="number" bind:value={startValue} />
                    <InputAddon>{currencySymbol}</InputAddon>
                </ButtonGroup>
            </Label>

            <Label>
                To
                <ButtonGroup>
                    <Input type="number" bind:value={endValue} />
                    <InputAddon>{currencySymbol}</InputAddon>
                </ButtonGroup>
            </Label>
        </div>

        <Heading tag="h5">Media</Heading>

        <Button color={alert.GifUrl ? "alternative" : "primary"} class="w-full" size="xl" outline on:click={() => { gifPickerOpen = true; }}>
            <ImageOutline class="w-5 h-5 me-1" />
            {#if alert.GifUrl} Change {:else} Pick or Upload {/if}
            GIF
        </Button>

        <Button color={alert.SoundUrl ? "alternative" : "primary"} class="w-full" size="xl" outline on:click={() => { soundPickerOpen = true; }}>
            <VolumeUpSolid class="w-5 h-5 me-1" />
            {#if alert.SoundUrl} Change {:else} Pick or Upload {/if}
            Sound
        </Button>

        <Heading tag="h5">Message</Heading>

        <Label>
            Template
            <Textarea bind:value={alert.Message} />
        </Label>

        <div class="flex flex-row space-x-2 m-auto justify-center">
            <RadioButton value={Alignment.START} bind:group={alert.Alignment}><BarsFromLeftSolid/></RadioButton>
            <RadioButton value={Alignment.CENTER} bind:group={alert.Alignment}><AlignCenterSolid/></RadioButton>
            <RadioButton value={Alignment.END} bind:group={alert.Alignment}><div class="transform rotate-180 scale-y-[-1]"><BarsFromLeftSolid/></div></RadioButton>
            <RadioButton value={Alignment.JUSTIFY} bind:group={alert.Alignment}><BarsSolid/></RadioButton>
        </div>

        <Heading tag="h6">Text Position</Heading>

        <div class="flex flex-col items-center space-y-2">
            <Radio value={Position.TOP} bind:group={alert.TextPosition}/>
            <div class="flex flex-row space-x-2">
                <Radio value={Position.LEFT} bind:group={alert.TextPosition}/>
                <div class="bg-gray-600 w-16 h-16 rounded-md text-center text-white flex items-center justify-center font-medium">
                    GIF
                </div>
                <Radio class="ms-1.5" value={Position.RIGHT} bind:group={alert.TextPosition}/>
            </div>
            <Radio value="{Position.BOTTOM}" bind:group={alert.TextPosition}/>
        </div>

        <Heading tag="h5">Look</Heading>

        <Label>
            Animation Type
            <Select bind:value={alert.Animation} items={animations}/>
        </Label>

        <Label>
            Animation Speed
            <Range min={Speed.OFF} max={Speed.FASTER} step="1" bind:value={alert.AnimationSpeed}/>
        </Label>

        <div class="flex space-x-10 pb-2">
            <Label class="flex flex-col">
                Text Color
                <input type="color" bind:value={alert.TextColor} />
            </Label>

            <Label class="flex flex-col">
                Accent Color
                <input type="color" bind:value={alert.AccentColor} />
            </Label>
        </div>

        <Button on:click={addNew} class="w-full">
            <PlusSolid class="w-5 h-5 me-1" />
            Add
        </Button>
    </div>

</Drawer>

<GifPicker bind:url={alert.GifUrl} searchTerm="money" bind:open={gifPickerOpen} />
<SoundPicker bind:url={alert.SoundUrl} bind:open={soundPickerOpen} />

{#if !hidden}
    <AlertViewer alerts={[alert]} event={exampleEvent} visible={!soundPickerOpen && !gifPickerOpen}/>
{/if}