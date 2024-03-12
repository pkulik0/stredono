<script lang="ts">
    import { saveAlert } from '$lib/alerts';
    import GifPicker from '$lib/comp/GifPicker.svelte';
    import SoundPicker from '$lib/comp/SoundPicker.svelte';
    import { settingsStore } from '$lib/events_settings';
    import { Alert, Alignment, AnimationType, Position, Speed } from '$lib/pb/alert_pb';
    import { Currency } from '$lib/pb/enums_pb';
    import { Event, EventType } from '$lib/pb/event_pb';
    import { userStore } from '$lib/user';
    import { pbEnumToItems } from '$lib/util';
    import {
        Button,
        ButtonGroup,
        Checkbox,
        CloseButton,
        Drawer,
        Heading,
        Input,
        InputAddon,
        Label,
        Radio,
        RadioButton,
        Range,
        Select,
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
    import { locale, t } from 'svelte-i18n';
    import { sineIn } from 'svelte/easing';
    import { fade } from 'svelte/transition';
    import AlertViewer from '$lib/comp/AlertViewer.svelte';

    export let hidden = true;
    export let eventType: EventType;

    export let alert = Alert.fromJson({
        ID: "",
        EventType: eventType,
        Min: 0,
        Max: 0,
        TextColor: "#FFFFFF",
        AccentColor: "#2381f8",
        Animation: AnimationType.PULSE,
        AnimationSpeed: Speed.MEDIUM,
        Alignment: Alignment.JUSTIFY,
        TextPosition: Position.BOTTOM,
    })
    $: if(alert.ID === "") {
        alert.EventType = eventType;
    }

    let startValue = "0";
    let endValue = "10";
    let hasMax = false;
    $: if(hasMax) {
        alert.Max = Number.parseFloat(endValue);
    } else {
        alert.Max = undefined;
    }

    let exampleEvent = new Event();
    $: {
        exampleEvent.Type = eventType;
        exampleEvent.Data = {
            "User": "John",
            "Value": "2137",
            "Currency": currencySymbol,
            "Message": "This is an example message. See which settings you like the most and adjust them to your needs."
        }
    }

    const addNew = async () => {
        try {
            await saveAlert(alert);
            hidden = true;
        } catch (e) {
            console.error(e); // TODO: Show error to user
        }
    }

    $: animations = pbEnumToItems(AnimationType);
    $: if ($locale) {
        animations = pbEnumToItems(AnimationType);
    }

    let gifPickerOpen = false;
    let soundPickerOpen = false;

    $: currency = $settingsStore?.Tips?.Currency;
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

    let transitionParams = {
        x: 300,
        duration: 200,
        easing: sineIn,
    };
    let divClass = "overflow-y-auto z-50 p-4 bg-gray-100 dark:bg-gray-800";
    let triggerClass = hasMax ? "" : "flex-row";

    const eventTypeNames = JSON.parse(JSON.stringify(EventType));

    const onShown = (e: Event) => {
    }
</script>

<Drawer activateClickOutside={false} transitionType="fly" {transitionParams} bind:hidden {divClass} placement="right" width="w-96">
    <div class="flex items-center">
        <h5 id="drawer-label" class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400">
            <BellActiveSolid class="w-4 h-4 me-2.5" />
            {$t("new_alert")} ({$t(eventTypeNames[eventType].toLowerCase())})
        </h5>
        <CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
    </div>

    <div class="w-full space-y-4">
        <Heading tag="h5">{$t("trigger")}</Heading>

        <div class="space-x-4 flex {triggerClass}">
            <Label class={hasMax ? "w-1/2" : "w-full"}>
                {$t("range_from")}
                <ButtonGroup class="w-full">
                    <Input type="number" bind:value={startValue} />
                    <InputAddon>{currencySymbol}</InputAddon>
                </ButtonGroup>
            </Label>

            {#if hasMax}
                <div class="w-1/2" transition:fade>
                    <Label>
                        {$t("range_to")}
                        <ButtonGroup class="w-full">
                            <Input type="number" bind:value={endValue} />
                            <InputAddon>{currencySymbol}</InputAddon>
                        </ButtonGroup>
                    </Label>
                </div>
            {/if}
        </div>
        <Checkbox bind:checked={hasMax} class="mt-1.5 ms-1">{$t("set_max")}</Checkbox>

        <Heading tag="h5">{$t("media")}</Heading>

        <Button color={alert.GifUrl ? "alternative" : "primary"} class="w-full" size="xl" outline on:click={() => { gifPickerOpen = true; }}>
            <ImageOutline class="w-5 h-5 me-1" />
            {#if alert.GifUrl} {$t("change")} {:else} {$t("pick_or_upload")} {/if}
            {$t("gif")}
        </Button>

        <Button color={alert.SoundUrl ? "alternative" : "primary"} class="w-full" size="xl" outline on:click={() => { soundPickerOpen = true; }}>
            <VolumeUpSolid class="w-5 h-5 me-1" />
            {#if alert.SoundUrl} {$t("change")} {:else} {$t("pick_or_upload")} {/if}
            {$t("sound")}
        </Button>

        <Heading tag="h5">{$t("appearance")}</Heading>

        <Label>
            {$t("animation_type")}
            <Select bind:value={alert.Animation} items={animations}/>
        </Label>

        <Label>
            {$t("animation_speed")}
            <Range min={Speed.OFF} max={Speed.FASTER} step="1" bind:value={alert.AnimationSpeed}/>
        </Label>

        <div class="flex space-x-10 pb-2">
            <Label class="flex flex-col">
                {$t("text_color")}
                <input type="color" bind:value={alert.TextColor} />
            </Label>

            <Label class="flex flex-col">
                {$t("accent_color")}
                <input type="color" bind:value={alert.AccentColor} />
            </Label>
        </div>

        <Heading tag="h6">{$t("text_position")}</Heading>

        <div class="flex flex-row space-x-2 pb-2 m-auto justify-center">
            <RadioButton value={Alignment.START} bind:group={alert.Alignment}><BarsFromLeftSolid/></RadioButton>
            <RadioButton value={Alignment.CENTER} bind:group={alert.Alignment}><AlignCenterSolid/></RadioButton>
            <RadioButton value={Alignment.END} bind:group={alert.Alignment}><div class="transform rotate-180 scale-y-[-1]"><BarsFromLeftSolid/></div></RadioButton>
            <RadioButton value={Alignment.JUSTIFY} bind:group={alert.Alignment}><BarsSolid/></RadioButton>
        </div>

        <div class="flex flex-col items-center space-y-2 pb-2">
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

        <Button on:click={addNew} class="w-full">
            <PlusSolid class="w-5 h-5 me-1" />
            {$t("save")}
        </Button>
    </div>

</Drawer>

<GifPicker bind:url={alert.GifUrl} searchTerm={$t("default_gif_search_term")} bind:open={gifPickerOpen} />
<SoundPicker bind:url={alert.SoundUrl} bind:open={soundPickerOpen} />

{#if !hidden}
    <AlertViewer isTest alerts={[alert]} event={exampleEvent} visible={!soundPickerOpen && !gifPickerOpen} {onShown}/>
{/if}