<script lang="ts">
    import {Button, CloseButton, Drawer, Fileupload, Hr, Input, Label, Select, Textarea,} from "flowbite-svelte";
    import {uploadToStorage} from "$lib/firebase/storage";
    import {BellActiveSolid, ImageOutline, PhoneOutline, PlusSolid} from "flowbite-svelte-icons";
    import {sineIn} from "svelte/easing";
    import {fade} from "svelte/transition";
    import {v4 as uuidv4} from "uuid";
    import ColorPicker, {ChromeVariant} from "svelte-awesome-color-picker";
    import {AlertStyle, AlertType, AnimationType, TriggerType} from "$lib/pb/user_pb";
    import {onMount} from "svelte";
    import GifPicker from "$lib/comp/GifPicker.svelte";
    import type {Gif} from "$lib/ext/tenor";
    import axios from "axios";

    const addNew = async () => {
        if(!gifFileList || !soundFileList || gifFileList.length === 0 || soundFileList.length === 0) {
            return;
        }

        const gifFile = gifFileList[0];
        const soundFile = soundFileList[0];

        try {
            const id = uuidv4().replace(/-/g, "");
            const gifUrl = await uploadToStorage("gifs", id, gifFile, false); // TODO: change name etc
            const soundUrl = await uploadToStorage("sounds", id, soundFile, false);

            hidden = true;
        } catch (e) {
            console.error(e); // TODO: Show error to user
        }
    }

    export let hidden = true;
    export let styles: Map<string, AlertStyle> = new Map();

    let soundFileList: FileList;

    let startValue = 0;
    let endValue = 10;

    let template = "[USER] donated [AMOUNT] [CURRENCY]!";

    let textColor = "#FFFFFF";
    let accentColor = "#e55430";

    let animations = [
        { "value": AnimationType.BOUNCE, "name": "Bounce" },
        { "value": AnimationType.SHAKE, "name": "Shake" },
    ];
    let animation = AnimationType.BOUNCE;

    let types = [
        { "value": AlertType.DONATE, "name": "Donation" },
        { "value": AlertType.SUBSCRIBE, "name": "Subscription" },
        { "value": AlertType.FOLLOW, "name": "Follow" },
        { "value": AlertType.BITS, "name": "Bits" },
        { "value": AlertType.RAID, "name": "Raid" },
    ];
    let type = AlertType.DONATE;

    let triggers = [
        { "value": TriggerType.AMOUNT,  "name": "Amount" },
        { "value": TriggerType.TIME,  "name": "Time" }
    ]
    let trigger = TriggerType.AMOUNT;


    interface Preset {
        value: AlertStyle
        name: string
    }
    let presets: Preset[] = [];
    let preset: AlertStyle|undefined = undefined;

    let gifPickerOpen = false;
    let gifFile: File|undefined = undefined;
    $: gifImage = gifFile ? URL.createObjectURL(gifFile) : null;
    $: console.log(gifImage);

    // Drawer settings
    let transitionParams = {
        x: 300,
        duration: 200,
        easing: sineIn,
    };
    let divClass = "overflow-y-auto z-50 p-4 rounded-xl m-0 md:m-4 bg-gray-100 dark:bg-gray-800";

    onMount(() => {
        styles.forEach((style, key) => {
            presets.push({ "value": style, "name": key });
        })
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


        <Label class="px-4">
            Template
            <Textarea bind:value={template} />
        </Label>

        <Hr/>

        <Label>
            Trigger
            <Select bind:value={trigger} items={triggers}/>
        </Label>

        <div class="space-x-4 flex flex-row px-4">
            <Label>
                From
                <Input type="number" bind:value={startValue} />
            </Label>


            <Label>
                To
                <Input type="number" bind:value={endValue} />
            </Label>
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

            <Label>
                GIF

                <div class="flex flex-col space-y-2">
                    <Button size="xl" outline on:click={() => { gifPickerOpen = true; }}>
                        <ImageOutline class="w-5 h-5 me-1" />
                        Pick or Upload GIF
                    </Button>
                </div>

                {#if gifImage}
                    <img src={gifImage} alt="GIF" class="px-8 py-2 rounded-lg w-full" in:fade />
                {/if}
            </Label>

            <Label>
                Sound
                <Fileupload size="sm" bind:files={soundFileList} accept="audio/*" />
            </Label>

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