<script lang="ts">
    import {
        Button, ButtonGroup,
        CloseButton,
        Drawer,
        Fileupload,
        Hr,
        Input,
        InputAddon,
        Label, P, Popover,
        Select,
        Textarea,
    } from "flowbite-svelte";
    import {uploadToStorage} from "$lib/firebase/storage";
    import {
        BellActiveSolid,
        ImageOutline,
        PhoneOutline,
        PlusSolid,
        QuestionCircleSolid,
        VolumeUpSolid
    } from "flowbite-svelte-icons";
    import {sineIn} from "svelte/easing";
    import {fade, slide} from "svelte/transition";
    import {v4 as uuidv4} from "uuid";
    import ColorPicker, {ChromeVariant} from "svelte-awesome-color-picker";
    import {
        AlertStyle,
        AlertType,
        AmountTrigger,
        AnimationType, Currency,
        TextToSpeechSettings
    } from "$lib/pb/user_pb";
    import {onMount} from "svelte";
    import GifPicker from "$lib/comp/GifPicker.svelte";
    import type {Gif} from "$lib/ext/tenor";
    import axios from "axios";
    import {Alert} from "$lib/pb/user_pb";
    import {userStore} from "$lib/user";
    import SoundPicker from "$lib/comp/SoundPicker.svelte";

    const getUuid = () => {
        return uuidv4().replace(/-/g, "");
    }

    const addNew = async () => {
        if (!gifFile) {
            console.error("GIF is missing");
            return;
        }
        if (!soundFile) {
            console.error("Sound is missing");
            return;
        }

        try {
            const id = getUuid();
            const gifUrl = await uploadToStorage("gifs", getUuid(), gifFile, false);
            const soundUrl = await uploadToStorage("sounds", getUuid(), soundFile, false);

            let alert = new Alert();
            alert.id = id;
            alert.type = type;
            alert.template = template;

            alert.style = new AlertStyle();
            alert.style.animation = animation;
            alert.style.gifUrl = gifUrl;
            alert.style.soundUrl = soundUrl;
            alert.style.textColor = textColor;
            alert.style.accentColor = accentColor;

            alert.amountTrigger = new AmountTrigger()
            alert.amountTrigger.min = startValue;
            alert.amountTrigger.max = endValue;

            alert.ttsSettings = new TextToSpeechSettings()
            hidden = true;
        } catch (e) {
            console.error(e); // TODO: Show error to user
        }
    }

    export let hidden = true;
    export let styles: Map<string, AlertStyle> = new Map();

    let startValue = 0;
    let endValue = 10;

    let intervalValue = 10;

    let template = "[USER] donated [AMOUNT] [CURRENCY]!";

    let textColor = "#FFFFFF";
    let accentColor = "#e55430";

    let animations = [
        { "value": AnimationType.BOUNCE, "name": "Bounce" },
        { "value": AnimationType.SHAKE, "name": "Shake" },
    ];
    let animation = AnimationType.BOUNCE;

    let types = [
        { "value": AlertType.DONATE, "name": "Donation / Cheer" },
        { "value": AlertType.SUBSCRIBE, "name": "Subscription" },
        { "value": AlertType.SUBGIFT, "name": "Gifted Subscription" },
        { "value": AlertType.FOLLOW, "name": "Follow" },
        { "value": AlertType.RAID, "name": "Raid" },
    ];
    let type = AlertType.DONATE;

    interface Preset {
        value: AlertStyle
        name: string
    }
    let presets: Preset[] = [];
    let preset: AlertStyle|undefined = undefined;

    let gifPickerOpen = false;
    let gifFile: File|undefined = undefined;
    $: gifImage = gifFile ? URL.createObjectURL(gifFile) : null;

    let soundPickerOpen = false;
    let soundFile: File|undefined = undefined;

    // Drawer settings
    let transitionParams = {
        x: 300,
        duration: 200,
        easing: sineIn,
    };
    let divClass = "overflow-y-auto z-50 p-4 rounded-xl m-0 md:m-4 bg-gray-100 dark:bg-gray-800";

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
        styles.forEach((style, key) => {
            presets.push({ "value": style, "name": key });
        })

        return userStore.subscribe((user) => {
            if (!user) {
                currency = undefined;
                return;
            }
            currency = user.currency;
        });
    })
</script>

<Drawer activateClickOutside={false} transitionType="fly" {transitionParams} bind:hidden {divClass} placement="right" width="w-96">
    <div class="flex items-center">
        <h5 id="drawer-label" class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400">
            <BellActiveSolid class="w-4 h-4 me-2.5" />
            New Alert
        </h5>
        <CloseButton on:click={() => (hidden = true)} class="mb-4 dark:text-white" />
    </div>

    <div class="w-full space-y-4">
        <Label>
            Type
            <Select bind:value={type} items={types}/>
        </Label>



        <div class="px-4 space-y-2">
            <Label>
                Template
                <Textarea bind:value={template} />
            </Label>

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

            {#if type === AlertType.DONATE}
                <div transition:slide>
                    <P weight="light" class="text-sm text-justify text-gray-500 dark:text-gray-400">
                        100 Bits are equal to 1 USD or its equivalent in other currencies.
                    </P>
                </div>
            {/if}
        </div>



        <Hr/>

        <Label>
            Preset
            <Select placeholder="None" bind:value={preset} items={presets}/>
        </Label>

        <div class="space-y-3 px-4">

            <Label>
                Animation
                <Select bind:value={animation} items={animations}/>
            </Label>

            <Button class="w-full" size="xl" outline on:click={() => { gifPickerOpen = true; }}>
                <ImageOutline class="w-5 h-5 me-1" />
                Pick or Upload GIF
            </Button>

            {#if gifImage}
                <img src={gifImage} alt="GIF" class="px-8 py-2 rounded-lg w-full" transition:fade />
            {/if}

            <Button class="w-full" size="xl" outline on:click={() => { soundPickerOpen = true; }}>
                <VolumeUpSolid class="w-5 h-5 me-1" />
                Pick or Upload Sound
            </Button>

            <div class="pt-2">
                <Label>
                    Text Color <br/>
                    <ColorPicker textInputModes={["hex"]} label="" isAlpha={false} isDialog sliderDirection="horizontal" components={ChromeVariant} bind:hex={textColor} />
                </Label>

                <Label>
                    Accent Color <br/>
                    <ColorPicker textInputModes={["hex"]} label="" isAlpha={false} isDialog sliderDirection="horizontal" components={ChromeVariant} bind:hex={accentColor} />
                </Label>
            </div>
        </div>

        <Button on:click={addNew} class="mt-7 w-full">
            <PlusSolid class="w-5 h-5 me-1" />
            Add
        </Button>
    </div>

</Drawer>

<GifPicker bind:file={gifFile} searchTerm="money" bind:open={gifPickerOpen} />
<SoundPicker bind:file={soundFile} bind:open={soundPickerOpen} />