import { useAccount } from '@/components/providers/account-provider';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Form, FormControl, FormField, FormItem } from '@/components/ui/form';
import { Separator } from '@/components/ui/separator';
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet';
import {
  isSystemTransformer,
  isUserDefinedTransformer,
  Transformer,
} from '@/shared/transformers';
import { getTransformerDataTypesString } from '@/util/util';
import {
  convertTransformerConfigSchemaToTransformerConfig,
  convertTransformerConfigToForm,
  JobMappingTransformerForm,
} from '@/yup-validations/jobs';
import { useMutation } from '@connectrpc/connect-query';
import { yupResolver } from '@hookform/resolvers/yup';
import { TransformerConfig } from '@neosync/sdk';
import { validateUserJavascriptCode } from '@neosync/sdk/connectquery';
import {
  Cross2Icon,
  EyeOpenIcon,
  MixerHorizontalIcon,
  Pencil1Icon,
} from '@radix-ui/react-icons';
import { ReactElement, useState } from 'react';
import { useForm } from 'react-hook-form';
import {
  EditJobMappingTransformerConfigFormContext,
  EditJobMappingTransformerConfigFormValues,
} from '../new/transformer/transform-form-validations';
import TransformerForm from '../new/transformer/TransformerForms/TransformerForm';

interface Props {
  transformer: Transformer;
  disabled: boolean;
  value: JobMappingTransformerForm;
  onSubmit(newValue: JobMappingTransformerForm): void;
}

// Note: this has issues with re-rendering due to being embedded within the tanstack table.
// This will cause the sheet to close when the user clicks back onto the page.
// This is partially solved by memoizing the tanstack columns, but any time the columns need to re-render, this sheet will close if it's open.
export default function EditTransformerOptions(props: Props): ReactElement {
  const { transformer, disabled, value, onSubmit } = props;
  const [isSheetOpen, setIsSheetOpen] = useState(false);
  return (
    <Sheet open={isSheetOpen} onOpenChange={setIsSheetOpen}>
      <SheetTrigger asChild>
        <Button
          variant="outline"
          size="sm"
          disabled={disabled}
          className="hidden h-[36px] lg:flex"
          type="button"
        >
          {isUserDefinedTransformer(transformer) ? (
            <EyeOpenIcon />
          ) : (
            <Pencil1Icon />
          )}
        </Button>
      </SheetTrigger>
      <SheetContent className="w-[800px]">
        <SheetHeader>
          <div className="flex flex-row w-full">
            <div className="flex flex-col space-y-2 w-full">
              <div className="flex flex-row justify-between items-center">
                <div className="flex flex-row gap-2">
                  <SheetTitle>{transformer.name}</SheetTitle>
                  <Badge variant="outline">
                    {getTransformerDataTypesString(transformer.dataTypes)}
                  </Badge>
                </div>
                <SheetClose>
                  <Cross2Icon className="h-4 w-4" />
                  <span className="sr-only">Close</span>
                </SheetClose>
              </div>
              <SheetDescription>{transformer?.description}</SheetDescription>
            </div>
          </div>
          <Separator />
        </SheetHeader>
        <div className="pt-8">
          {transformer && (
            <EditTransformerConfig
              value={
                isSystemTransformer(transformer)
                  ? convertTransformerConfigSchemaToTransformerConfig(
                      value.config
                    )
                  : (transformer.config ?? new TransformerConfig())
              }
              onSubmit={(newval) => {
                onSubmit({
                  ...value,
                  config: convertTransformerConfigToForm(newval),
                });
                setIsSheetOpen(false);
              }}
              isDisabled={disabled || isUserDefinedTransformer(transformer)}
            />
          )}
        </div>
      </SheetContent>
    </Sheet>
  );
}

interface EditTransformerConfigProps {
  value: TransformerConfig;
  onSubmit(newValue: TransformerConfig): void;
  isDisabled: boolean;
}

function EditTransformerConfig(
  props: EditTransformerConfigProps
): ReactElement {
  const { value, onSubmit, isDisabled } = props;

  const { account } = useAccount();
  const { mutateAsync: isJavascriptCodeValid } = useMutation(
    validateUserJavascriptCode
  );

  const form = useForm<
    EditJobMappingTransformerConfigFormValues,
    EditJobMappingTransformerConfigFormContext
  >({
    mode: 'onChange',
    resolver: yupResolver(EditJobMappingTransformerConfigFormValues),
    defaultValues: { config: convertTransformerConfigToForm(value) },
    context: {
      accountId: account?.id ?? '',
      isUserJavascriptCodeValid: isJavascriptCodeValid,
    },
  });
  return (
    <Form {...form}>
      <div className="flex flex-col gap-8">
        <FormField
          control={form.control}
          name="config"
          render={({ field }) => (
            <FormItem>
              <FormControl>
                <TransformerForm
                  value={convertTransformerConfigSchemaToTransformerConfig(
                    field.value
                  )}
                  setValue={(newValue) => {
                    field.onChange(convertTransformerConfigToForm(newValue));
                  }}
                  disabled={isDisabled}
                  errors={form.formState.errors}
                  NoConfigComponent={<NoAdditionalTransformerConfigurations />}
                />
              </FormControl>
            </FormItem>
          )}
        />
        <div className="flex justify-end">
          <Button
            type="button"
            disabled={isDisabled || !form.formState.isValid}
            onClick={(e) => {
              form.handleSubmit((formValues) => {
                onSubmit(
                  convertTransformerConfigSchemaToTransformerConfig(
                    formValues.config
                  )
                );
              })(e);
            }}
          >
            Save
          </Button>
        </div>
      </div>
    </Form>
  );
}

function NoAdditionalTransformerConfigurations(): ReactElement {
  return (
    <Alert className="border-gray-200 dark:border-gray-700 shadow-sm">
      <div className="flex flex-row items-center gap-4">
        <MixerHorizontalIcon className="h-4 w-4" />
        <AlertDescription className="text-gray-500">
          There are no additional configurations for this Transformer
        </AlertDescription>
      </div>
    </Alert>
  );
}
