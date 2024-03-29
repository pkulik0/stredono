<script lang="ts">
    import { getDefaultAlert, saveAlert } from '$lib/alerts';
    import EventViewer from '$lib/comp/EventViewer.svelte';
    import GifPicker from '$lib/comp/GifPicker.svelte';
    import SoundPicker from '$lib/comp/SoundPicker.svelte';
    import { Alert, Alignment, AnimationType, Position, Speed } from '$lib/pb/alert_pb';
    import { Currency } from '$lib/pb/enums_pb';
    import { Event, EventType } from '$lib/pb/event_pb';
    import { currencyToSymbol, pbEnumToItems } from '$lib/util';
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
        Select
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
    import { alertStore } from './stores';

    export let hidden = true;
    export let eventType: EventType;

    $: alert = $alertStore
    $: if(alert.ID === "") {
        alert.EventType = eventType;
    }

    let minValue = "1";
    $: alert.Min = alert.EventType == EventType.TIP ? Number.parseFloat(minValue) : Number.parseInt(minValue)

    let maxValue = "10";
    $: alert.Max = hasMax ? (alert.EventType == EventType.TIP ? Number.parseFloat(maxValue) : Number.parseInt(maxValue)) : undefined

    let hasMax = false;

    $: exampleEvent = Event.fromJson({
        ID: "",
        Type: eventType,
        Data: {
            "User": $t("example_user"),
            "Value": $t("example_value"),
            "Currency": currencySymbol,
            "Message": $t("example_message"),
        }
    })

    const addNew = async () => {
        try {
            await saveAlert(alert);
            onClose()
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

    $: currencySymbol = currencyToSymbol(Currency.PLN)

    let transitionParams = {
        x: 300,
        duration: 200,
        easing: sineIn,
    };
    let divClass = "overflow-y-auto z-50 p-4 bg-gray-100 dark:bg-gray-800";
    let triggerClass = hasMax ? "" : "flex-row";

    const eventTypeNames = JSON.parse(JSON.stringify(EventType));

    $: valueDescription = (() => {
        switch(alert.EventType) {
            case EventType.TIP:
                return currencySymbol;
            case EventType.CHEER:
                return $t("bits_count");
            case EventType.SUB:
                return $t("months");
            case EventType.SUB_GIFT:
                return $t("count");
            case EventType.RAID:
                return $t("viewers_count");
            default:
                return "";
        }
    })();

    const onClose = () => {
        hidden = true;
        alertStore.set(getDefaultAlert(eventType));
    }
</script>

<Drawer activateClickOutside={false} transitionType="fly" {transitionParams} bind:hidden {divClass} placement="right" width="w-96">
    <div class="flex items-center">
        <h5 id="drawer-label" class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400">
            <BellActiveSolid class="w-4 h-4 me-2.5" />
            {$t("new_alert")} ({$t(eventTypeNames[eventType].toLowerCase())})
        </h5>
        <CloseButton on:click={onClose} class="mb-4 dark:text-white" />
    </div>

    <div class="w-full space-y-4">
        {#if valueDescription}
            <Heading tag="h5">{$t("trigger")}</Heading>

            <div class="space-x-4 flex {triggerClass}">
                <Label class={hasMax ? "w-1/2" : "w-full"}>
                    {$t("range_from")}
                    <ButtonGroup class="w-full">
                        <Input type="number" bind:value={minValue} min="1" max={hasMax ? maxValue : undefined}/>
                        <InputAddon>{valueDescription}</InputAddon>
                    </ButtonGroup>
                </Label>

                {#if hasMax}
                    <div class="w-1/2" transition:fade>
                        <Label>
                            {$t("range_to")}
                            <ButtonGroup class="w-full">
                                <Input type="number" bind:value={maxValue} min={minValue}/>
                                <InputAddon>{valueDescription}</InputAddon>
                            </ButtonGroup>
                        </Label>
                    </div>
                {/if}
            </div>
            <Checkbox bind:checked={hasMax} class="mt-1.5 ms-1">{$t("set_max")}</Checkbox>
        {/if}

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

{#if !hidden && (!soundPickerOpen && !gifPickerOpen)}
    <EventViewer alerts={[alert]} events={[exampleEvent]} isTest/>
{/if}